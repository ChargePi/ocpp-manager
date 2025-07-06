package controllers

import (
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
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (r *ReservationCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (r *ReservationCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewReservationCtrlr() *ReservationCtrlr {
	return &ReservationCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  []variables.VariableName{}, // No required variables for ReservationCtrlr
		supportedVariables: supportedReservationVariables(),
	}
}
