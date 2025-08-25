package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameDateTime                         variables.VariableName = "DateTime"
	VariableNameNtpSource                        variables.VariableName = "NtpSource"
	VariableNameNtpServerUri                     variables.VariableName = "NtpServerUri"
	VariableNameTimeOffset                       variables.VariableName = "TimeOffset"
	VariableNameNextTimeOffsetTransitionDateTime variables.VariableName = "NextTimeOffsetTransitionDateTime"
	VariableNameTimeOffsetNextTransition         variables.VariableName = "TimeOffsetNextTransition"
	VariableNameTimeSource                       variables.VariableName = "TimeSource"
	VariableNameTimeZone                         variables.VariableName = "TimeZone"
	VariableNameTimeAdjustmentReportingThreshold variables.VariableName = "TimeAdjustmentReportingThreshold"
)

func requiredClockVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameDateTime,
		VariableNameTimeSource,
	}
}

func optionalClockVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameTimeZone,
		VariableNameNtpSource,
		VariableNameNtpServerUri,
		VariableNameTimeOffsetNextTransition,
		VariableNameNextTimeOffsetTransitionDateTime,
		VariableNameTimeOffset,
		VariableNameTimeAdjustmentReportingThreshold,
	}
}

// supportedClockVariables returns a list of all variables supported by the ClockCtrlr.
func supportedClockVariables() []variables.VariableName {
	return append(requiredClockVariables(), optionalClockVariables()...)
}

type ClockCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (c *ClockCtrlr) GetName() component.Name {
	return component.ComponentNameClockCtrlr
}

func (c *ClockCtrlr) GetInstanceId() string {
	return c.instanceId
}

func (c *ClockCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (c *ClockCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (c *ClockCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (c *ClockCtrlr) GetRequiredVariables() []variables.VariableName {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.requiredVariables
}

func (c *ClockCtrlr) GetSupportedVariables() []variables.VariableName {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.supportedVariables
}

func (c *ClockCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !c.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	variable, exists := c.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (c *ClockCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
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

func (c *ClockCtrlr) Validate(key variables.VariableName) bool {
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

func NewClockCtrlr() *ClockCtrlr {
	ctrlr := &ClockCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredClockVariables(),
		supportedVariables: supportedClockVariables(),
		instanceId:         "clock-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
