package prompt

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInput(t *testing.T) {
	tests := []struct {
		v string
		e bool
	}{
		{v: "hello\n", e: false},
		{v: "hello", e: true},
	}
	for _, test := range tests {
		input = strings.NewReader(test.v)
		output, err := GetInput()
		if err != nil {
			assert.Equal(t, test.e, true)
			continue
		}
		assert.Equal(t, test.v, output)
	}
}
