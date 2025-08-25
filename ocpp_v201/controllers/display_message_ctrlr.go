package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameDisplayMessageEnabled             variables.VariableName = "Enabled"
	VariableNameDisplayMessageAvailable           variables.VariableName = "Available"
	VariableNameNumberOfDisplayMessages           variables.VariableName = "DisplayMessages"
	VariableNameDisplayMessageSupportedFormats    variables.VariableName = "SupportedFormats"
	VariableNameDisplayMessageSupportedPriorities variables.VariableName = "SupportedPriorities"
)

func requiredDisplayMessageVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameNumberOfDisplayMessages,
		VariableNameDisplayMessageSupportedFormats,
		VariableNameDisplayMessageSupportedPriorities,
	}
}

func optionalDisplayMessageVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameDisplayMessageEnabled,
		VariableNameDisplayMessageAvailable,
	}
}

func supportedDisplayMessageVariables() []variables.VariableName {
	return append(requiredDisplayMessageVariables(), optionalDisplayMessageVariables()...)
}

type DisplayCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (d *DisplayCtrlr) GetName() component.Name {
	return component.ComponentNameDisplayMessageCtrlr
}

func (d *DisplayCtrlr) GetInstanceId() string {
	return d.instanceId
}

func (d *DisplayCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (d *DisplayCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (d *DisplayCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (d *DisplayCtrlr) GetRequiredVariables() []variables.VariableName {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.requiredVariables
}

func (d *DisplayCtrlr) GetSupportedVariables() []variables.VariableName {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.supportedVariables
}

func (d *DisplayCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !d.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	variable, exists := d.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}

	return variable, nil
}

func (d *DisplayCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !d.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	v, exists := d.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (d *DisplayCtrlr) Validate(key variables.VariableName) bool {
	if !d.validator.IsVariableSupported(key) {
		return false
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	v, exists := d.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewDisplayCtrlr() *DisplayCtrlr {
	ctrlr := &DisplayCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredDisplayMessageVariables(),
		supportedVariables: supportedDisplayMessageVariables(),
		instanceId:         "display-message-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
