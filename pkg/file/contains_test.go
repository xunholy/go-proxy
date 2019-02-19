package file

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		v []string
		f string
		e bool
	}{
		{v: []string{"start", "all"}, f: "all", e: true},
		{v: []string{"start", ""}, f: "all", e: false},
	}

	for _, test := range tests {
		result := Contains(test.v, test.f)
		assert.Equal(t, test.e, result)
	}
}
