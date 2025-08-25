package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameBytesPerMessage        variables.VariableName = "BytesPerMessage"
	VariableNameConfigurationValueSize variables.VariableName = "ConfigurationValueSize"
	VariableNameReportingValueSize     variables.VariableName = "ReportingValueSize"
	VariableNameItemsPerMessage        variables.VariableName = "ItemsPerMessage"
)

func requiredVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameBytesPerMessage,
		VariableNameConfigurationValueSize,
		VariableNameReportingValueSize,
		VariableNameItemsPerMessage,
	}
}

func optionalVariables() []variables.VariableName {
	return []variables.VariableName{}
}

func supportedVariables() []variables.VariableName {
	return append(requiredVariables(), optionalVariables()...)
}

type DeviceDataCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (d *DeviceDataCtrlr) GetName() component.ComponentName {
	return component.ComponentNameDeviceDataCtrlr
}

func (d *DeviceDataCtrlr) GetInstanceId() string {
	return d.instanceId
}

func (d *DeviceDataCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (d *DeviceDataCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (d *DeviceDataCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (d *DeviceDataCtrlr) GetRequiredVariables() []variables.VariableName {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.requiredVariables
}

func (d *DeviceDataCtrlr) GetSupportedVariables() []variables.VariableName {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.supportedVariables
}

func (d *DeviceDataCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !d.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	variable, exists := d.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (d *DeviceDataCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !d.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	v, exists := d.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (d *DeviceDataCtrlr) Validate(key variables.VariableName) bool {
	if !d.validator.IsVariableSupported(key) {
		return false
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	// Check if the variable exists
	v, exists := d.variables[key]
	if !exists {
		return false
	}

	// Validate the variable itself
	return v.Validate()
}

// ValidateAllRequiredVariables checks that all required variables are present and valid
func (d *DeviceDataCtrlr) ValidateAllRequiredVariables() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, requiredVar := range d.requiredVariables {
		// Check if the variable exists
		v, exists := d.variables[requiredVar]
		if !exists {
			return false
		}
		// Validate the variable itself
		if !v.Validate() {
			return false
		}
	}

	return true
}

func NewDeviceDataCtrlr() *DeviceDataCtrlr {
	ctrlr := &DeviceDataCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredVariables(),
		supportedVariables: supportedVariables(),
		instanceId:         "device-data-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	// Initialize all required variables with default values
	ctrlr.variables[VariableNameBytesPerMessage] = variables.NewVariable(
		VariableNameBytesPerMessage,
		variables.VariableTypeInteger,
	)
	ctrlr.variables[VariableNameConfigurationValueSize] = variables.NewVariable(
		VariableNameConfigurationValueSize,
		variables.VariableTypeInteger,
	)
	ctrlr.variables[VariableNameReportingValueSize] = variables.NewVariable(
		VariableNameReportingValueSize,
		variables.VariableTypeInteger,
	)
	ctrlr.variables[VariableNameItemsPerMessage] = variables.NewVariable(
		VariableNameItemsPerMessage,
		variables.VariableTypeInteger,
	)

	return ctrlr
}
