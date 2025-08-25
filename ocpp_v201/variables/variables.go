package variables

import "errors"

type Mutability string

const (
	MutabilityReadOnly  Mutability = "ReadOnly"
	MutabilityReadWrite Mutability = "ReadWrite"
	MutabilityWriteOnly Mutability = "WriteOnly"
)

type VariableType string

const (
	VariableTypeString       VariableType = "string"
	VariableTypeInteger      VariableType = "integer"
	VariableTypeNumber       VariableType = "number"
	VariableTypeBool         VariableType = "boolean"
	VariableTypeOptionList   VariableType = "OptionList"
	VariableTypeSequenceList VariableType = "SequenceList"
	VariableTypeMemberList   VariableType = "MemberList"
)

type VariableName string

type Variable struct {
	Name VariableName
	// attributes are conditionally mutable.
	attributes map[string]VariableAttributes
	// Characteristics are read-only
	Characteristics VariableCharacteristic
}

// Validate checks if all variable attributes are valid.
func (v *Variable) Validate() bool {
	for _, attributes := range v.attributes {
		if attributes.Validate() == false {
			return false
		}
	}

	// todo validate according to characteristics

	return true
}

// UpdateVariableAttribute updates the variable attribute if it exists and if the value is valid
func (v *Variable) UpdateVariableAttribute(attribute string, value interface{}) error {
	// Check if exists
	existingEntry, ok := v.attributes[attribute]
	if !ok {
		return errors.New("Variable attribute not found: " + attribute)
	}

	// Check if it we can even update it
	if existingEntry.Mutability == MutabilityReadOnly {
		return errors.New("Variable attribute is read-only: " + attribute)
	}

	// Check if the operation is allowed
	existingEntry.Value = value
	if !existingEntry.Validate() {
		return errors.New("invalid value")
	}

	// Update the value
	v.attributes[attribute] = existingEntry
	return nil
}

// GetVariableAttribute gets the variable attribute if it exists.
func (v *Variable) GetVariableAttribute(attribute string) (*VariableAttributes, error) {
	// Check if exists
	existingEntry, ok := v.attributes[attribute]
	if !ok {
		return nil, errors.New("Variable attribute not found: " + attribute)
	}

	// Check if it is readable
	if existingEntry.Mutability == MutabilityWriteOnly {
		return nil, errors.New("Variable attribute is write-only: " + attribute)
	}

	return &existingEntry, nil
}

// GetAllAttributes returns a copy of all attributes for this variable.
func (v *Variable) GetAllAttributes() map[string]VariableAttributes {
	result := make(map[string]VariableAttributes, len(v.attributes))
	for k, vAttr := range v.attributes {
		result[k] = vAttr
	}
	return result
}

type VariableAttributes struct {
	Type       VariableType
	Mutability Mutability
	Value      interface{}
}

// Validate validates
func (va *VariableAttributes) Validate() bool {
	if va == nil {
		return false
	}

	switch va.Mutability {
	case MutabilityReadOnly, MutabilityReadWrite, MutabilityWriteOnly:
	default:
		return false
	}

	switch va.Type {
	case VariableTypeNumber:
		// Must be castable to float
		_, castable := va.Value.(float64)
		return castable
	case VariableTypeBool:
		// Must be castable to bool
		_, castable := va.Value.(bool)
		return castable
	case VariableTypeInteger:
		// Must be castable to bool
		_, castable := va.Value.(int64)
		return castable
	case VariableTypeOptionList, VariableTypeMemberList, VariableTypeSequenceList:
		_, castable := va.Value.([]interface{})
		return castable
	case VariableTypeString:
		_, castable := va.Value.(string)
		return castable
	}

	return false
}

func (va *VariableAttributes) copy() VariableAttributes {
	return *va
}

// Update updates the variable attribute value
func (va *VariableAttributes) Update(value interface{}) error {
	// Make a copy of the variable attributes to validate the new value
	attrsCopy := va.copy()
	attrsCopy.Value = value
	if !attrsCopy.Validate() {
		return errors.New("invalid value for variable attribute")
	}

	va.Value = value
	return nil
}

type VariableCharacteristic struct {
	DataType   string
	MaxLimit   *int
	MinLimit   *int
	Unit       *string
	ValuesList []string
}

// NewVariable creates a new variable with the given name, type, and default value
func NewVariable(name VariableName, varType VariableType) *Variable {
	return &Variable{
		Name:       name,
		attributes: map[string]VariableAttributes{},
	}
}
