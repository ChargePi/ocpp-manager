package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name     string
		opts     []GetSetVariableOption
		expected componentVariableOptions
	}{
		{
			name: "default options",
			expected: componentVariableOptions{
				attributeType: "",
			},
			opts: []GetSetVariableOption{},
		},
		{
			name: "with attribute type",
			expected: componentVariableOptions{
				attributeType: "abc",
			},
			opts: []GetSetVariableOption{
				WithAttributeType("abc"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := componentVariableOptions{}
			for _, opt := range tt.opts {
				opt(&opts)
			}

			assert.Equal(t, tt.expected, opts)
		})
	}
}
