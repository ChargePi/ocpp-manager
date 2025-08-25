package component

type GetSetVariableOption func(o *componentVariableOptions)

type componentVariableOptions struct {
	attributeType string
}

// WithAttributeType sets the attribute type for the variable options.
func WithAttributeType(attributeType string) GetSetVariableOption {
	return func(o *componentVariableOptions) {
		o.attributeType = attributeType
	}
}
