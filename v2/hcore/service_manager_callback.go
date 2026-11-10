package hcore

import (
	"github.com/bobthebuilder1234232243/donutvpn-core/v2/service_manager"
	"github.com/sagernet/sing-box/adapter"
)

type donutvpnMainServiceManager struct{}

var _ adapter.LifecycleService = (*donutvpnMainServiceManager)(nil)

func (h *donutvpnMainServiceManager) Name() string { return "donutvpnMainServiceManager" }
func (h *donutvpnMainServiceManager) Start(stage adapter.StartStage, _ *adapter.Scope) error {
	if stage == adapter.StartStateStarted {
		return service_manager.OnMainServiceStart()
	}
	return nil
}

func (h *donutvpnMainServiceManager) Close() error {
	return service_manager.OnMainServiceClose()
}
