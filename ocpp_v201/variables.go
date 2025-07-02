package ocpp_v201

type Mutability string

const (
	MutabilityReadOnly  Mutability = "ReadOnly"
	MutabilityReadWrite Mutability = "ReadWrite"
	MutabilityWriteOnly Mutability = "WriteOnly"
)

type VariableName string

type Variable struct {
	Name            VariableName
	Attributes      VariableAttributes
	Characteristics VariableCharacteristic
	Value           interface{}
}

type VariableAttributes struct {
	Mutability Mutability
}

type VariableCharacteristic struct {
	DataType   string
	MaxLimit   *int
	Unit       *string
	ValuesList []string
}
