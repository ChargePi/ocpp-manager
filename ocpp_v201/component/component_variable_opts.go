package component

type GetSetVariableOption func(o *componentVariableOptions)

type componentVariableOptions struct {
	attributeType string
}

func WithAttributeType(attributeType string) GetSetVariableOption {
	return func(o *componentVariableOptions) {
		o.attributeType = attributeType
	}
}
