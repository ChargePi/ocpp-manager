package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameLocalAuthListEnabled                variables.VariableName = "Enabled"
	VariableNameLocalAuthListEntries                variables.VariableName = "Entries"
	VariableNameLocalAuthListItemsPerMessage        variables.VariableName = "ItemsPerMessage"
	VariableNameLocalAuthListBytesPerMessage        variables.VariableName = "BytesPerMessage"
	VariableNameLocalAuthListStorage                variables.VariableName = "Storage"
	VariableNameLocalAuthListDisablePostAuthorize   variables.VariableName = "DisablePostAuthorize"
	VariableNameLocalAuthListSupportsExpiryDateTime variables.VariableName = "SupportsExpiryDateTime"
)

func requiredLocalAuthListVariables() []variables.VariableName {
	return []variables.VariableName{}
}

func optionalLocalAuthListVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameLocalAuthListItemsPerMessage,
		VariableNameLocalAuthListBytesPerMessage,
		VariableNameLocalAuthListStorage,
		VariableNameLocalAuthListDisablePostAuthorize,
		VariableNameLocalAuthListSupportsExpiryDateTime,
		VariableNameLocalAuthListEnabled,
		VariableNameLocalAuthListEntries,
	}
}

func supportedLocalAuthListVariables() []variables.VariableName {
	return append(requiredLocalAuthListVariables(), optionalLocalAuthListVariables()...)
}

type LocalAuthListCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (l *LocalAuthListCtrlr) GetName() component.Name {
	return component.ComponentNameLocalAuthListCtrlr
}

func (l *LocalAuthListCtrlr) GetInstanceId() string {
	return l.instanceId
}

func (l *LocalAuthListCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (l *LocalAuthListCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (l *LocalAuthListCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (l *LocalAuthListCtrlr) GetRequiredVariables() []variables.VariableName {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.requiredVariables
}

func (l *LocalAuthListCtrlr) GetSupportedVariables() []variables.VariableName {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.supportedVariables
}

func (l *LocalAuthListCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !l.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	variable, exists := l.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (l *LocalAuthListCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !l.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	v, exists := l.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (l *LocalAuthListCtrlr) Validate(key variables.VariableName) bool {
	if !l.validator.IsVariableSupported(key) {
		return false
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	v, exists := l.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewLocalAuthListCtrlr() *LocalAuthListCtrlr {
	ctrlr := &LocalAuthListCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredLocalAuthListVariables(),
		supportedVariables: supportedLocalAuthListVariables(),
		instanceId:         "local-auth-list-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
