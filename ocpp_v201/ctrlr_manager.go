package ocpp_v201

import (
	"errors"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/controllers"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
	"sync"
)

// Global manager instance
var (
	managerInstance *Manager
	managerOnce     sync.Once
)

func init() {
	managerOnce.Do(func() {
		managerInstance = NewManager()
	})
}

func GetManager() *Manager {
	return managerInstance
}

type Manager struct {
	controllers map[component.ComponentName]component.Component
}

type ManagerOption func(*managerOptions)

type managerOptions struct {
	// Supported profiles + controllers
}

func NewManager() *Manager {
	return &Manager{
		controllers: map[component.ComponentName]component.Component{
			component.ComponentNameMonitoringCtrlr:     controllers.NewMonitoringCtrlr(),
			component.ComponentNameLocalAuthListCtrlr:  controllers.NewLocalAuthListCtrlr(),
			component.ComponentNameReservationCtrlr:    controllers.NewReservationCtrlr(),
			component.ComponentNameSmartChargingCtrlr:  controllers.NewSmartChargingCtrlr(),
			component.ComponentNameTxCtrlr:             controllers.NewTxCtrlr(),
			component.ComponentNameSecurityCtrlr:       controllers.NewSecurityCtrlr(),
			component.ComponentNameClockCtrlr:          controllers.NewClockCtrlr(),
			component.ComponentNameDeviceDataCtrlr:     controllers.NewDeviceDataCtrlr(),
			component.ComponentNameAuthCtrlr:           controllers.NewAuthCtrlr(),
			component.ComponentNameAuthCacheCtrlr:      controllers.NewAuthCacheCtrlr(),
			component.ComponentNameISO15118Ctrlr:       controllers.NewISO15118Ctrlr(),
			component.ComponentNameDisplayMessageCtrlr: controllers.NewDisplayCtrlr(),
			component.ComponentNameOCPPCommCtrlr:       controllers.NewOCPPCommCtrlr(),
			component.ComponentNameAlignedDataCtrlr:    controllers.NewAlignedDataCtrlr(),
			component.ComponentNameSampledDataCtrlr:    controllers.NewSampledDataCtrlr(),
			component.ComponentNameTariffCostCtrlr:     controllers.NewTariffCostCtrlr(),
			component.ComponentNameCustomizationCtrlr:  controllers.NewCustomizationCtrlr(),
		},
	}
}

func (m *Manager) UpdateVariable(name component.ComponentName, variableName variables.VariableName, attributeName string, attributeValue interface{}) error {
	controller, ok := m.controllers[name]
	if !ok {
		return errors.New("controller not found")
	}

	return controller.UpdateVariable(variableName, attributeName, attributeValue)
}

func (m *Manager) GetVariable(name component.ComponentName, variableName variables.VariableName) (*variables.Variable, error) {
	controller, ok := m.controllers[name]
	if !ok {
		return nil, errors.New("controller not found")
	}

	return controller.GetVariable(variableName)
}
