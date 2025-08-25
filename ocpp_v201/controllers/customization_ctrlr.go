package controllers

import (
	"fmt"
	"slices"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameCustomImplementationEnabled variables.VariableName = "CustomImplementationEnabled"
)

type CustomizationCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName // Set
	supportedVariables []variables.VariableName // Set
	instanceId         string
	validator          *variableValidator
}

func (c *CustomizationCtrlr) GetName() component.ComponentName {
	return component.ComponentNameCustomizationCtrlr
}

func (c *CustomizationCtrlr) GetInstanceId() string {
	return c.instanceId
}

func (c *CustomizationCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (c *CustomizationCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (c *CustomizationCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (c *CustomizationCtrlr) GetRequiredVariables() []variables.VariableName {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.requiredVariables
}

func (c *CustomizationCtrlr) GetSupportedVariables() []variables.VariableName {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.supportedVariables
}

func (c *CustomizationCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !c.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	variable, exists := c.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return &variable, nil
}

func (c *CustomizationCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !c.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	v, exists := c.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (c *CustomizationCtrlr) Validate(key variables.VariableName) bool {
	if !c.validator.IsVariableSupported(key) {
		return false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	v, exists := c.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func (c *CustomizationCtrlr) AddRequiredVariable(key variables.VariableName) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if slices.Contains(c.requiredVariables, key) || slices.Contains(c.supportedVariables, key) {
		return
	}

	c.requiredVariables = append(c.requiredVariables, key)
}

func (c *CustomizationCtrlr) AddSupportedVariable(key variables.VariableName) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if slices.Contains(c.supportedVariables, key) || slices.Contains(c.requiredVariables, key) {
		return
	}

	c.supportedVariables = append(c.supportedVariables, key)
}

func NewCustomizationCtrlr() *CustomizationCtrlr {
	ctrlr := &CustomizationCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  make([]variables.VariableName, 0),
		supportedVariables: make([]variables.VariableName, 0),
		instanceId:         "customization-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
