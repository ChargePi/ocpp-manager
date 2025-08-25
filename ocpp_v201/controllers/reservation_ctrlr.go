package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameReservationEnabled         variables.VariableName = "Enabled"
	VariableNameReservationAvailable       variables.VariableName = "Available"
	VariableNameReservationNonEvseSpecific variables.VariableName = "NonEvseSpecific"
)

func optionalReservationVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameReservationEnabled,
		VariableNameReservationAvailable,
		VariableNameReservationNonEvseSpecific,
	}
}

func supportedReservationVariables() []variables.VariableName {
	return optionalReservationVariables()
}

type ReservationCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (r *ReservationCtrlr) GetName() component.Name {
	return component.ComponentNameReservationCtrlr
}

func (r *ReservationCtrlr) GetInstanceId() string {
	return r.instanceId
}

func (r *ReservationCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (r *ReservationCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (r *ReservationCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (r *ReservationCtrlr) GetRequiredVariables() []variables.VariableName {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.requiredVariables
}

func (r *ReservationCtrlr) GetSupportedVariables() []variables.VariableName {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.supportedVariables
}

func (r *ReservationCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !r.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	variable, exists := r.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (r *ReservationCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !r.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	v, exists := r.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (r *ReservationCtrlr) Validate(key variables.VariableName) bool {
	if !r.validator.IsVariableSupported(key) {
		return false
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	v, exists := r.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewReservationCtrlr() *ReservationCtrlr {
	ctrlr := &ReservationCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  []variables.VariableName{}, // No required variables for ReservationCtrlr
		supportedVariables: supportedReservationVariables(),
		instanceId:         "reservation-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)
	return ctrlr
}
