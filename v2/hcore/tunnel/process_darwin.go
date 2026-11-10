package tunnel

import (
	"github.com/bobthebuilder1234232243/donutvpn-core/v2/hcore/adminbg"
	"github.com/sagernet/sing-box/experimental/libbox"
)

func FindConnectionOwner(ipProtocol int32, sourceAddress string, sourcePort int32) (*libbox.ConnectionOwner, error) {
	return adminbg.FindConnectionOwner(ipProtocol, sourceAddress, sourcePort)
}
