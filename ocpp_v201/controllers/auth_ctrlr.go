package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameAuthEnabled                   variables.VariableName = "Enabled"
	VariableNameAdditionalInfoItemsPerMessage variables.VariableName = "AdditionalInfoItemsPerMessage"
	VariableNameOfflineTxForUnknownIdEnabled  variables.VariableName = "OfflineTxForUnknownIdEnabled"
	VariableNameAuthorizeRemoteStart          variables.VariableName = "AuthorizeRemoteStart"
	VariableNameLocalAuthorizeOffline         variables.VariableName = "LocalAuthorizeOffline"
	VariableNameLocalPreAuthorize             variables.VariableName = "LocalPreAuthorize"
	VariableNameMasterPassGroupId             variables.VariableName = "MasterPassGroupId"
	VariableNameDisableRemoteAuthorization    variables.VariableName = "DisableRemoteAuthorization"
)

func requiredAuthVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAuthorizeRemoteStart,
		VariableNameLocalAuthorizeOffline,
		VariableNameLocalPreAuthorize,
	}
}

func optionalAuthVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAuthEnabled,
		VariableNameAdditionalInfoItemsPerMessage,
		VariableNameOfflineTxForUnknownIdEnabled,
		VariableNameMasterPassGroupId,
		VariableNameDisableRemoteAuthorization,
	}
}

func supportedAuthVariables() []variables.VariableName {
	return append(requiredAuthVariables(), optionalAuthVariables()...)
}

type AuthCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (a *AuthCtrlr) GetName() component.Name {
	return component.ComponentNameAuthCtrlr
}

func (a *AuthCtrlr) GetInstanceId() string {
	return a.instanceId
}

func (a *AuthCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (a *AuthCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (a *AuthCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (a *AuthCtrlr) GetRequiredVariables() []variables.VariableName {
	return a.requiredVariables
}

func (a *AuthCtrlr) GetSupportedVariables() []variables.VariableName {
	return a.supportedVariables
}

func (a *AuthCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
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

func (a *AuthCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
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

func (a *AuthCtrlr) Validate(key variables.VariableName) bool {
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

func NewAuthCtrlr() *AuthCtrlr {
	ctrlr := &AuthCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredAuthVariables(),
		supportedVariables: supportedAuthVariables(),
		instanceId:         "auth-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
