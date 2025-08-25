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
	components map[component.ComponentName]component.Component
}

type ManagerOption func(*managerOptions)

type managerOptions struct {
	// Supported profiles + components
	components []component.Component
}

func WithComponents(components []component.Component) ManagerOption {
	return func(opts *managerOptions) {
		opts.components = components
	}
}

func NewManager(opts ...ManagerOption) *Manager {
	manager := &Manager{
		components: make(map[component.ComponentName]component.Component),
	}

	defaults := managerOptions{
		components: []component.Component{
			controllers.NewMonitoringCtrlr(),
			controllers.NewLocalAuthListCtrlr(),
			controllers.NewReservationCtrlr(),
			controllers.NewSmartChargingCtrlr(),
			controllers.NewTxCtrlr(),
			controllers.NewSecurityCtrlr(),
			controllers.NewClockCtrlr(),
			controllers.NewDeviceDataCtrlr(),
			controllers.NewAuthCtrlr(),
			controllers.NewAuthCacheCtrlr(),
			controllers.NewISO15118Ctrlr(),
			controllers.NewDisplayCtrlr(),
			controllers.NewOCPPCommCtrlr(),
			controllers.NewAlignedDataCtrlr(),
			controllers.NewSampledDataCtrlr(),
			controllers.NewTariffCostCtrlr(),
			controllers.NewCustomizationCtrlr(),
		},
	}
	for _, opt := range opts {
		opt(&defaults)
	}

	// Register components
	for _, component := range defaults.components {
		manager.components[component.GetName()] = component
	}

	return manager
}

func (m *Manager) UpdateVariable(name component.ComponentName, variableName variables.VariableName, attributeName string, attributeValue interface{}) error {
	controller, ok := m.components[name]
	if !ok {
		return errors.New("controller not found")
	}

	return controller.UpdateVariable(variableName, attributeName, attributeValue)
}

func (m *Manager) GetVariable(name component.ComponentName, variableName variables.VariableName) (*variables.Variable, error) {
	controller, ok := m.components[name]
	if !ok {
		return nil, errors.New("controller not found")
	}

	return controller.GetVariable(variableName)
}
