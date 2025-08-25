package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameMonitoringEnabled                      variables.VariableName = "Enabled"
	VariableNameMonitoringAvailable                    variables.VariableName = "Available"
	VariableNameItemsPerMessageClearVariableMonitoring variables.VariableName = "ItemsPerMessage"
	VariableNameItemsPerMessageSetVariableMonitoring   variables.VariableName = "ItemsPerMessage"
	VariableNameClearVariableMonitoring                variables.VariableName = "BytesPerMessage"
	VariableNameBytesPerMessageSetVariableMonitoring   variables.VariableName = "BytesPerMessage"
	VariableNameOfflineMonitoringEventQueuingSeverity  variables.VariableName = "OfflineQueuingSeverity"
	VariableNameActiveMonitoringBase                   variables.VariableName = "ActiveMonitoringBase"
	VariableNameActiveMonitoringLevel                  variables.VariableName = "ActiveMonitoringLevel"
)

func requiredMonitoringVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameItemsPerMessageSetVariableMonitoring,
		VariableNameBytesPerMessageSetVariableMonitoring,
	}
}

func optionalMonitoringVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameClearVariableMonitoring,
		VariableNameMonitoringEnabled,
		VariableNameMonitoringAvailable,
		VariableNameItemsPerMessageClearVariableMonitoring,
		VariableNameOfflineMonitoringEventQueuingSeverity,
		VariableNameActiveMonitoringBase,
		VariableNameActiveMonitoringLevel,
	}
}

func supportedMonitoringVariables() []variables.VariableName {
	return append(requiredMonitoringVariables(), optionalMonitoringVariables()...)
}

type MonitoringCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (m *MonitoringCtrlr) GetName() component.Name {
	return component.ComponentNameMonitoringCtrlr
}

func (m *MonitoringCtrlr) GetInstanceId() string {
	return m.instanceId
}

func (m *MonitoringCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (m *MonitoringCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (m *MonitoringCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (m *MonitoringCtrlr) GetRequiredVariables() []variables.VariableName {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.requiredVariables
}

func (m *MonitoringCtrlr) GetSupportedVariables() []variables.VariableName {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.supportedVariables
}

func (m *MonitoringCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !m.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	variable, exists := m.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}

	return variable, nil
}

func (m *MonitoringCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !m.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	v, exists := m.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (m *MonitoringCtrlr) Validate(key variables.VariableName) bool {
	if !m.validator.IsVariableSupported(key) {
		return false
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	v, exists := m.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewMonitoringCtrlr() *MonitoringCtrlr {
	ctrlr := &MonitoringCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredMonitoringVariables(),
		supportedVariables: supportedMonitoringVariables(),
		instanceId:         "monitoring-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)
	return ctrlr
}
