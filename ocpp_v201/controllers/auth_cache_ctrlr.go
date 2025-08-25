package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameAuthCacheEnabled              variables.VariableName = "Enabled"
	VariableNameAuthCacheAvailable            variables.VariableName = "Available"
	VariableNameAuthCacheLifeTime             variables.VariableName = "LifeTime"
	VariableNameAuthCacheStorage              variables.VariableName = "Storage"
	VariableNameAuthCachePolicy               variables.VariableName = "Policy"
	VariableNameAuthCacheDisablePostAuthorize variables.VariableName = "DisablePostAuthorize"
)

func requiredAuthCacheVariables() []variables.VariableName {
	return []variables.VariableName{}
}

func optionalAuthCacheVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAuthCacheAvailable,
		VariableNameAuthCacheEnabled,
		VariableNameAuthCacheLifeTime,
		VariableNameAuthCacheStorage,
		VariableNameAuthCachePolicy,
		VariableNameAuthCacheDisablePostAuthorize,
	}
}

func supportedAuthCacheVariables() []variables.VariableName {
	return append(requiredAuthCacheVariables(), optionalAuthCacheVariables()...)
}

type AuthCacheCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	supportedVariables []variables.VariableName
	requiredVariables  []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (a *AuthCacheCtrlr) GetName() component.Name {
	return component.ComponentNameAuthCacheCtrlr
}

func (a *AuthCacheCtrlr) GetInstanceId() string {
	return a.instanceId
}

func (a *AuthCacheCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (a *AuthCacheCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (a *AuthCacheCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (a *AuthCacheCtrlr) GetRequiredVariables() []variables.VariableName {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.requiredVariables
}

func (a *AuthCacheCtrlr) GetSupportedVariables() []variables.VariableName {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.supportedVariables
}

func (a *AuthCacheCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !a.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	variable, exists := a.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (a *AuthCacheCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !a.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	v, exists := a.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (a *AuthCacheCtrlr) Validate(key variables.VariableName) bool {
	if !a.validator.IsVariableSupported(key) {
		return false
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	v, exists := a.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewAuthCacheCtrlr() *AuthCacheCtrlr {
	ctrlr := &AuthCacheCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		supportedVariables: supportedAuthCacheVariables(),
		requiredVariables:  requiredAuthCacheVariables(),
		instanceId:         "auth-cache-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)
	return ctrlr
}
