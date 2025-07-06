package controllers

import variables2 "github.com/ChargePi/ocpp-manager/ocpp_v201/variables"

type ChademoCtrlr struct {
	variables          map[variables2.VariableName]variables2.Variable
	requiredVariables  []variables2.VariableName
	supportedVariables []variables2.VariableName
}

func NewChademoCtrlr() *ChademoCtrlr {
	return &ChademoCtrlr{}
}
