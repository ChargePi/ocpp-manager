package controllers

import variables "github.com/ChargePi/ocpp-manager/ocpp_v201/variables"

type ChademoCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func NewChademoCtrlr() *ChademoCtrlr {
	return &ChademoCtrlr{}
}
