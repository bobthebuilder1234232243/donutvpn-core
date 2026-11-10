package hcore

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/bobthebuilder1234232243/donutvpn-core/v2/config"
	"golang.org/x/net/proxy"

	"github.com/sagernet/sing-box/option"
)

func getRandomAvailblePort() uint16 {
	// TODO: implement it
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	return uint16(listener.Addr().(*net.TCPAddr).Port)
}

func RunInstanceString(ctx context.Context, donutvpnSettings *config.DonutOptions, proxiesInput string) (*DonutInstance, error) {
	if donutvpnSettings == nil {
		donutvpnSettings = config.DefaultDonutOptions()
	}

	singconfigs, err := config.ParseConfig(ctx, &config.ReadOptions{Content: proxiesInput}, true, donutvpnSettings, false)
	if err != nil {
		return nil, err
	}
	return RunInstance(ctx, donutvpnSettings, singconfigs)
}

func RunInstance(ctx context.Context, donutvpnSettings *config.DonutOptions, singconfig *option.Options) (*DonutInstance, error) {
	if donutvpnSettings == nil {
		donutvpnSettings = config.DefaultDonutOptions()
	}
	donutvpnSettings.EnableClashApi = false
	donutvpnSettings.InboundOptions.MixedPort = getRandomAvailblePort()
	donutvpnSettings.InboundOptions.EnableTun = false
	donutvpnSettings.InboundOptions.EnableTunService = false
	donutvpnSettings.InboundOptions.SetSystemProxy = false
	donutvpnSettings.InboundOptions.TProxyPort = 0
	donutvpnSettings.InboundOptions.DirectPort = 0
	donutvpnSettings.InboundOptions.RedirectPort = 0
	donutvpnSettings.Region = "other"
	donutvpnSettings.BlockAds = false
	donutvpnSettings.LogFile = "/dev/null"

	finalConfigs, err := config.BuildConfig(ctx, donutvpnSettings, &config.ReadOptions{Options: singconfig})
	if err != nil {
		return nil, err
	}

	instance, err := NewService(ctx, *finalConfigs, nil)
	if err != nil {
		return nil, err
	}

	<-time.After(250 * time.Millisecond)
	hservice := &DonutInstance{
		StartedService: instance,
		ListenPort:     donutvpnSettings.InboundOptions.MixedPort}
	hservice.PingCloudflare()
	return hservice, nil
}

// dialer, err := s.libbox.GetInstance().Router().Dialer(context.Background())

func (s *DonutInstance) Close() error {
	return closeService(s.StartedService)
}

func (s *DonutInstance) GetContent(url string) (string, error) {
	return s.ContentFromURL("GET", url, 10*time.Second)
}

func (s *DonutInstance) ContentFromURL(method string, url string, timeout time.Duration) (string, error) {
	if method == "" {
		return "", fmt.Errorf("empty method")
	}
	if url == "" {
		return "", fmt.Errorf("empty url")
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	dialer, err := proxy.SOCKS5("tcp", fmt.Sprintf("127.0.0.1:%d", s.ListenPort), nil, proxy.Direct)
	if err != nil {
		return "", err
	}

	transport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return "", fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if body == nil {
		return "", fmt.Errorf("empty body")
	}

	return string(body), nil
}

func (s *DonutInstance) PingCloudflare() (time.Duration, error) {
	return s.Ping("http://cp.cloudflare.com")
}

// func (s *DonutVpnService) RawConnection(ctx context.Context, url string) (net.Conn, error) {
// 	return
// }

func (s *DonutInstance) PingAverage(url string, count int) (time.Duration, error) {
	if count <= 0 {
		return -1, fmt.Errorf("count must be greater than 0")
	}

	var sum int
	real_count := 0
	for i := 0; i < count; i++ {
		delay, err := s.Ping(url)
		if err == nil {
			real_count++
			sum += int(delay.Milliseconds())
		} else if real_count == 0 && i > count/2 {
			return -1, fmt.Errorf("ping average failed")
		}

	}
	return time.Duration(sum / real_count * int(time.Millisecond)), nil
}

func (s *DonutInstance) Ping(url string) (time.Duration, error) {
	startTime := time.Now()
	_, err := s.ContentFromURL("HEAD", url, 4*time.Second)
	if err != nil {
		return -1, err
	}
	duration := time.Since(startTime)
	return duration, nil
}
