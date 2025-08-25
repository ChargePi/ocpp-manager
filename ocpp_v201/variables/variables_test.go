package variables

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
)

type variablesTestSuite struct {
	suite.Suite
}

func (suite *variablesTestSuite) TestVariable_Validate() {
	tests := []struct {
		name     string
		variable *Variable
		expected bool
	}{
		{
			name: "valid variable",
			variable: &Variable{
				Name: "TestVariable",
				attributes: map[string]VariableAttributes{
					"value": {
						Type:       VariableTypeString,
						Mutability: MutabilityReadWrite,
						Value:      "test",
					},
				},
				Characteristics: VariableCharacteristic{
					DataType: "string",
				},
			},
			expected: true,
		},
		{
			name: "variable with invalid attribute",
			variable: &Variable{
				Name: "TestVariable",
				attributes: map[string]VariableAttributes{
					"value": {
						Type:       VariableTypeString,
						Mutability: "InvalidMutability", // Invalid mutability
						Value:      "test",
					},
				},
				Characteristics: VariableCharacteristic{
					DataType: "string",
				},
			},
			expected: false,
		},
		{
			name: "variable with empty attributes",
			variable: &Variable{
				Name:       "TestVariable",
				attributes: map[string]VariableAttributes{},
				Characteristics: VariableCharacteristic{
					DataType: "string",
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Assert().Equal(tt.expected, tt.variable.Validate())
		})
	}
}

func (suite *variablesTestSuite) TestVariable_GetVariableAttribute() {
	var1 := &Variable{
		Name: "TestVariable",
		attributes: map[string]VariableAttributes{
			"readWrite": {
				Type:       VariableTypeString,
				Mutability: MutabilityReadWrite,
				Value:      "test",
			},
			"readOnly": {
				Type:       VariableTypeInteger,
				Mutability: MutabilityReadOnly,
				Value:      int64(42),
			},
			"writeOnly": {
				Type:       VariableTypeBool,
				Mutability: MutabilityWriteOnly,
				Value:      true,
			},
		},
	}

	tests := []struct {
		name          string
		attribute     string
		expectError   bool
		errorMsg      string
		expectedValue interface{}
	}{
		{
			name:          "get read-write attribute",
			attribute:     "readWrite",
			expectError:   false,
			expectedValue: "test",
		},
		{
			name:          "get read-only attribute",
			attribute:     "readOnly",
			expectError:   false,
			expectedValue: int64(42),
		},
		{
			name:        "get write-only attribute",
			attribute:   "writeOnly",
			expectError: true,
			errorMsg:    "Variable attribute is write-only: writeOnly",
		},
		{
			name:        "get non-existent attribute",
			attribute:   "nonExistent",
			expectError: true,
			errorMsg:    "Variable attribute not found: nonExistent",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			attr, err := var1.GetVariableAttribute(tt.attribute)
			if tt.expectError {
				suite.Assert().Error(err)
				suite.Assert().Equal(tt.errorMsg, err.Error())
				suite.Assert().Nil(attr)
			} else {
				suite.Require().NoError(err)
				suite.Assert().Equal(tt.expectedValue, attr.Value)
			}
		})
	}
}

func (suite *variablesTestSuite) TestVariable_UpdateVariableAttribute() {
	var1 := &Variable{
		Name: "TestVariable",
		attributes: map[string]VariableAttributes{
			"readWrite": {
				Type:       VariableTypeString,
				Mutability: MutabilityReadWrite,
				Value:      "old",
			},
			"readOnly": {
				Type:       VariableTypeInteger,
				Mutability: MutabilityReadOnly,
				Value:      int64(42),
			},
			"writeOnly": {
				Type:       VariableTypeBool,
				Mutability: MutabilityWriteOnly,
				Value:      false,
			},
		},
	}

	tests := []struct {
		name        string
		attribute   string
		value       interface{}
		expectError bool
		errorMsg    string
	}{
		{
			name:        "update read-write attribute with valid value",
			attribute:   "readWrite",
			value:       "new",
			expectError: false,
		},
		{
			name:        "update read-only attribute",
			attribute:   "readOnly",
			value:       int64(100),
			expectError: true,
			errorMsg:    "Variable attribute is read-only: readOnly",
		},
		{
			name:        "update write-only attribute with valid value",
			attribute:   "writeOnly",
			value:       true,
			expectError: false,
		},
		{
			name:        "update non-existent attribute",
			attribute:   "nonExistent",
			value:       "value",
			expectError: true,
			errorMsg:    "Variable attribute not found: nonExistent",
		},
		{
			name:        "update with invalid value type",
			attribute:   "readWrite",
			value:       123, // int instead of string
			expectError: true,
			errorMsg:    "invalid value",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := var1.UpdateVariableAttribute(tt.attribute, tt.value)
			if tt.expectError {
				suite.Assert().Error(err)
				suite.Assert().Equal(tt.errorMsg, err.Error())
			} else {
				suite.Require().NoError(err)
				// Verify the value was updated for readable attributes
				if tt.attribute == "readWrite" {
					attr, err := var1.GetVariableAttribute(tt.attribute)
					suite.Require().NoError(err)
					suite.Assert().Equal(tt.value, attr.Value)
				}
			}
		})
	}
}

func (suite *variablesTestSuite) TestVariableAttributes_Validate() {
	tests := []struct {
		name     string
		attr     *VariableAttributes
		expected bool
	}{
		{
			name: "valid string attribute",
			attr: &VariableAttributes{
				Type:       VariableTypeString,
				Mutability: MutabilityReadWrite,
				Value:      "test",
			},
			expected: true,
		},
		{
			name: "valid integer attribute",
			attr: &VariableAttributes{
				Type:       VariableTypeInteger,
				Mutability: MutabilityReadOnly,
				Value:      int64(42),
			},
			expected: true,
		},
		{
			name: "valid number attribute",
			attr: &VariableAttributes{
				Type:       VariableTypeNumber,
				Mutability: MutabilityReadWrite,
				Value:      float64(3.14),
			},
			expected: true,
		},
		{
			name: "valid boolean attribute",
			attr: &VariableAttributes{
				Type:       VariableTypeBool,
				Mutability: MutabilityWriteOnly,
				Value:      true,
			},
			expected: true,
		},
		{
			name: "valid list attributes",
			attr: &VariableAttributes{
				Type:       VariableTypeOptionList,
				Mutability: MutabilityReadWrite,
				Value:      []interface{}{"item1", "item2"},
			},
			expected: true,
		},
		{
			name:     "nil attribute",
			attr:     nil,
			expected: false,
		},
		{
			name: "invalid mutability",
			attr: &VariableAttributes{
				Type:       VariableTypeString,
				Mutability: "Invalid",
				Value:      "test",
			},
			expected: false,
		},
		{
			name: "invalid type for string",
			attr: &VariableAttributes{
				Type:       VariableTypeString,
				Mutability: MutabilityReadWrite,
				Value:      123, // int instead of string
			},
			expected: false,
		},
		{
			name: "invalid type for integer",
			attr: &VariableAttributes{
				Type:       VariableTypeInteger,
				Mutability: MutabilityReadWrite,
				Value:      "not an int", // string instead of int64
			},
			expected: false,
		},
		{
			name: "invalid type for number",
			attr: &VariableAttributes{
				Type:       VariableTypeNumber,
				Mutability: MutabilityReadWrite,
				Value:      "not a number", // string instead of float64
			},
			expected: false,
		},
		{
			name: "invalid type for boolean",
			attr: &VariableAttributes{
				Type:       VariableTypeBool,
				Mutability: MutabilityReadWrite,
				Value:      "not a bool", // string instead of bool
			},
			expected: false,
		},
		{
			name: "invalid type for list",
			attr: &VariableAttributes{
				Type:       VariableTypeOptionList,
				Mutability: MutabilityReadWrite,
				Value:      "not a list", // string instead of []interface{}
			},
			expected: false,
		},
		{
			name: "unknown variable type",
			attr: &VariableAttributes{
				Type:       "UnknownType",
				Mutability: MutabilityReadWrite,
				Value:      "test",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Assert().Equal(tt.expected, tt.attr.Validate())
		})
	}
}

func (suite *variablesTestSuite) TestVariableAttributes_Update() {
	tests := []struct {
		name          string
		attr          *VariableAttributes
		newValue      interface{}
		expectError   bool
		errorMsg      string
		expectedValue interface{}
	}{
		{
			name: "valid string update",
			attr: &VariableAttributes{
				Type:       VariableTypeString,
				Mutability: MutabilityReadWrite,
				Value:      "old",
			},
			newValue:      "new",
			expectError:   false,
			expectedValue: "new",
		},
		{
			name: "invalid string update",
			attr: &VariableAttributes{
				Type:       VariableTypeString,
				Mutability: MutabilityReadWrite,
				Value:      "old",
			},
			newValue:      123, // int instead of string
			expectError:   true,
			errorMsg:      "invalid value for variable attribute",
			expectedValue: "old", // Value should remain unchanged
		},
		{
			name: "valid integer update",
			attr: &VariableAttributes{
				Type:       VariableTypeInteger,
				Mutability: MutabilityReadWrite,
				Value:      int64(10),
			},
			newValue:      int64(20),
			expectError:   false,
			expectedValue: int64(20),
		},
		{
			name: "invalid integer update",
			attr: &VariableAttributes{
				Type:       VariableTypeInteger,
				Mutability: MutabilityReadWrite,
				Value:      int64(10),
			},
			newValue:      "not an int",
			expectError:   true,
			errorMsg:      "invalid value for variable attribute",
			expectedValue: int64(10), // Value should remain unchanged
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := tt.attr.Update(tt.newValue)
			if tt.expectError {
				suite.Assert().Error(err)
				suite.Assert().Equal(tt.errorMsg, err.Error())
			} else {
				suite.Require().NoError(err)
			}
			suite.Assert().Equal(tt.expectedValue, tt.attr.Value)
		})
	}
}

func (suite *variablesTestSuite) TestVariableAttributes_Copy() {
	original := &VariableAttributes{
		Type:       VariableTypeString,
		Mutability: MutabilityReadWrite,
		Value:      "original",
	}

	copy := original.copy()
	suite.Assert().Equal(original.Type, copy.Type)
	suite.Assert().Equal(original.Mutability, copy.Mutability)
	suite.Assert().Equal(original.Value, copy.Value)

	// Verify it's a deep copy
	copy.Value = "modified"
	suite.Assert().Equal("original", original.Value)
	suite.Assert().Equal("modified", copy.Value)
}

func (suite *variablesTestSuite) TestVariableCharacteristic() {
	// Test VariableCharacteristic struct
	char := VariableCharacteristic{
		DataType:   "string",
		MaxLimit:   lo.ToPtr(100),
		MinLimit:   lo.ToPtr(0),
		Unit:       lo.ToPtr("watts"),
		ValuesList: []string{"option1", "option2"},
	}

	suite.Assert().Equal("string", char.DataType)
	suite.Assert().Equal(100, *char.MaxLimit)
	suite.Assert().Equal(0, *char.MinLimit)
	suite.Assert().Equal("watts", *char.Unit)
	suite.Assert().Equal([]string{"option1", "option2"}, char.ValuesList)
}

func TestVariables(t *testing.T) {
	suite.Run(t, new(variablesTestSuite))
}
