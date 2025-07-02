package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
	"slices"
)

type variableValidator struct {
	component component.Component
}

func newVariableValidator(component component.Component) *variableValidator {
	return &variableValidator{
		component: component,
	}
}

// IsVariableSupported checks if the given variable name is supported by the component.
func (vv *variableValidator) IsVariableSupported(variableName variables.VariableName) bool {
	supportedVariables := vv.component.GetSupportedVariables()
	return slices.Contains(supportedVariables, variableName)
}
