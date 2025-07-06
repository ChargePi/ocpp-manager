package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameEVConnectionTimeOut      variables.VariableName = "EVConnectionTimeOut"
	VariableNameStopTxOnEVSideDisconnect variables.VariableName = "StopTxOnEVSideDisconnect"
	VariableNameTxBeforeAcceptedEnabled  variables.VariableName = "TxBeforeAcceptedEnabled"
	VariableNameTxStartPoint             variables.VariableName = "TxStartPoint"
	VariableNameTxStopPoint              variables.VariableName = "TxStopPoint"
	VariableNameMaxEnergyOnInvalidId     variables.VariableName = "MaxEnergyOnInvalidId"
	VariableNameStopTxOnInvalidId        variables.VariableName = "StopTxOnInvalidId"
)

func requiredTxCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameEVConnectionTimeOut,
		VariableNameStopTxOnEVSideDisconnect,
		VariableNameStopTxOnInvalidId,
		VariableNameTxStartPoint,
		VariableNameTxStopPoint,
	}
}

func optionalTxCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameTxBeforeAcceptedEnabled,
		VariableNameMaxEnergyOnInvalidId,
	}
}

func supportedTxCtrlrVariables() []variables.VariableName {
	return append(requiredTxCtrlrVariables(), optionalTxCtrlrVariables()...)
}

type TxCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (t *TxCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (t *TxCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewTxCtrlr() *TxCtrlr {
	return &TxCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredTxCtrlrVariables(),
		supportedVariables: supportedTxCtrlrVariables(),
	}
}
