package ocpp_v201

import (
	"errors"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/controllers"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
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

func RegisterComponent(component component.Component) error {
	return managerInstance.RegisterComponent(component)
}

type Manager struct {
	mu         sync.RWMutex
	components map[component.Name]component.Component
}

func NewManager(opts ...ManagerOption) *Manager {
	manager := &Manager{
		components: make(map[component.Name]component.Component),
	}

	defaults := managerOptions{
		// By default, instantiate all controllers
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
	for _, c := range defaults.components {
		err := manager.RegisterComponent(c)
		if err != nil {
			return nil
		}
	}

	return manager
}

func (m *Manager) UpdateVariable(name component.Name, variableName variables.VariableName, attributeName string, attributeValue interface{}) error {
	m.mu.RLock()
	controller, ok := m.components[name]
	m.mu.RUnlock()
	if !ok {
		return errors.New("controller not found")
	}

	return controller.UpdateVariable(variableName, attributeName, attributeValue)
}

func (m *Manager) GetVariable(name component.Name, variableName variables.VariableName) (*variables.Variable, error) {
	m.mu.RLock()
	controller, ok := m.components[name]
	m.mu.RUnlock()
	if !ok {
		return nil, errors.New("controller not found")
	}

	return controller.GetVariable(variableName)
}

func (m *Manager) RegisterComponent(component component.Component) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, found := m.components[component.GetName()]; !found {
		m.components[component.GetName()] = component
		return nil
	}

	return errors.New("component already registered")
}
