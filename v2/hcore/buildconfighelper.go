package hcore

import (
	"context"
	"encoding/json"
	"os"

	"github.com/bobthebuilder1234232243/donutvpn-core/v2/config"
	"github.com/bobthebuilder1234232243/donutvpn-core/v2/db"
	hcommon "github.com/bobthebuilder1234232243/donutvpn-core/v2/hcommon"
	hutils "github.com/bobthebuilder1234232243/donutvpn-core/v2/hutils"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/experimental/libbox"
	"github.com/sagernet/sing-box/option"
)

func BuildConfigJson(ctx context.Context, in *StartRequest) (string, error) {
	Log(LogLevel_DEBUG, LogType_CORE, "Stating Service ")

	parsedContent, err := BuildConfig(ctx, in)
	if err != nil {
		return "", err
	}
	res, err := parsedContent.MarshalJSONContext(ctx)
	return string(res), err
}

func BuildConfig(ctx context.Context, in *StartRequest) (*option.Options, error) {
	Log(LogLevel_DEBUG, LogType_CORE, "Building Config...")

	readOpt := &config.ReadOptions{Content: in.ConfigContent, Path: in.ConfigPath}
	if !in.EnableRawConfig {
		// hcontent, err := json.MarshalIndent(static.DonutOptions, "", " ")
		// if err != nil {
		// 	return nil, err
		// }

		// Log(LogLevel_DEBUG, LogType_CORE, "Building config ", string(hcontent))
		// Log(LogLevel_DEBUG, LogType_CORE, "Building config ")
		return config.BuildConfig(ctx, static.DonutOptions, readOpt)
	}
	return config.ReadSingOptions(ctx, readOpt)

}

func (s *CoreService) Parse(ctx context.Context, in *ParseRequest) (*ParseResponse, error) {
	return Parse(libbox.FromContext(ctx, nil), in)
}

func Parse(ctx context.Context, in *ParseRequest) (*ParseResponse, error) {
	defer config.DeferPanicToError("parse", func(err error) {
		Log(LogLevel_FATAL, LogType_CONFIG, err.Error())
		StopAndAlert(MessageType_UNEXPECTED_ERROR, err.Error())
	})

	path := in.TempPath
	if path == "" {
		path = in.ConfigPath
	}

	config, err := config.ParseConfigBytes(ctx, &config.ReadOptions{Content: in.Content, Path: path}, true, static.DonutOptions, false)
	if err != nil {
		return &ParseResponse{
			ResponseCode: hcommon.ResponseCode_FAILED,
			Message:      err.Error(),
		}, err
	}
	if in.ConfigPath != "" {
		err = os.WriteFile(in.ConfigPath, config, 0o644)
		if err != nil {
			return &ParseResponse{
				ResponseCode: hcommon.ResponseCode_FAILED,
				Message:      err.Error(),
			}, err
		}
	}
	return &ParseResponse{
		ResponseCode: hcommon.ResponseCode_OK,
		Content:      string(config),
		Message:      "",
	}, err
}

func (s *CoreService) ChangeDonutSettings(ctx context.Context, in *ChangeDonutSettingsRequest) (*CoreInfoResponse, error) {
	return ChangeDonutSettings(in, true)
}

func ChangeDonutSettings(in *ChangeDonutSettingsRequest, insert bool) (*CoreInfoResponse, error) {
	static.DonutOptions = config.DefaultDonutOptions()
	defer func() {
		switch static.DonutOptions.LogLevel {
		case "debug":
			static.logLevel.Store(int32(LogLevel_DEBUG))
		case "info":
			static.logLevel.Store(int32(LogLevel_INFO))
		case "warn":
			static.logLevel.Store(int32(LogLevel_WARNING))
		case "error":
			static.logLevel.Store(int32(LogLevel_ERROR))
		case "fatal":
			static.logLevel.Store(int32(LogLevel_FATAL))
		case "trace":
			static.logLevel.Store(int32(LogLevel_TRACE))
		default:
			static.logLevel.Store(int32(LogLevel_INFO))
		}
		static.debug = static.debug || LogLevel(static.logLevel.Load()) <= LogLevel_DEBUG
	}()

	if in.DonutSettingsJson == "" {
		return &CoreInfoResponse{}, nil
	}
	if insert {
		settings := db.GetTable[hcommon.AppSettings]()
		settings.UpdateInsert(&hcommon.AppSettings{
			Id:    "DonutSettingsJson",
			Value: in.DonutSettingsJson,
		})
	}

	err := json.Unmarshal([]byte(in.DonutSettingsJson), static.DonutOptions)
	if err != nil {
		return nil, err
	}

	if static.DonutOptions.Warp.WireguardConfigStr != "" {
		err := json.Unmarshal([]byte(static.DonutOptions.Warp.WireguardConfigStr), &static.DonutOptions.Warp.WireguardConfig)
		if err != nil {
			return nil, err
		}
	}
	if static.DonutOptions.Warp2.WireguardConfigStr != "" {
		err := json.Unmarshal([]byte(static.DonutOptions.Warp2.WireguardConfigStr), &static.DonutOptions.Warp2.WireguardConfig)
		if err != nil {
			return nil, err
		}
	}
	return &CoreInfoResponse{}, nil
}

func (s *CoreService) GenerateConfig(ctx context.Context, in *GenerateConfigRequest) (*GenerateConfigResponse, error) {
	return GenerateConfig(libbox.FromContext(ctx, nil), in)
}

func GenerateConfig(ctx context.Context, in *GenerateConfigRequest) (*GenerateConfigResponse, error) {
	defer config.DeferPanicToError("generateConfig", func(err error) {
		Log(LogLevel_FATAL, LogType_CONFIG, err.Error())
		StopAndAlert(MessageType_UNEXPECTED_ERROR, err.Error())
	})
	if static.DonutOptions == nil {
		static.DonutOptions = config.DefaultDonutOptions()
	}
	config, err := config.ParseBuildConfigBytes(ctx, static.DonutOptions, &config.ReadOptions{Path: in.Path})
	if err != nil {
		return nil, err
	}

	return &GenerateConfigResponse{
		ConfigContent: string(config),
	}, nil
}

func removeTunnelIfNeeded(options *option.Options) (tuninb *option.TunInboundOptions) {
	if hutils.TunAllowed() {
		return nil
	}

	// Create a new slice to hold the remaining inbounds
	newInbounds := make([]option.Inbound, 0, len(options.Inbounds))

	for _, inb := range options.Inbounds {
		if inb.Type == C.TypeTun {
			if d, ok := inb.Options.(option.TunInboundOptions); ok {
				tuninb = &d
			}

		} else {
			newInbounds = append(newInbounds, inb)
		}
	}

	options.Inbounds = newInbounds
	return tuninb
}
