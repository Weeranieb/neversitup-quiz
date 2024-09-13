package utils_test

import (
	"neverquiz/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmileyFace(t *testing.T) {
	tests := []struct {
		input  []string
		output int
	}{
		{[]string{":)", ";(", ";}", ":-D"}, 2},
		{[]string{";D", ":-(", ":-)", ";~)"}, 3},
		{[]string{";]", ":[", ";*", ":$", ";-D"}, 1},
	}

	for _, test := range tests {
		if got := utils.CountSmileyFace(test.input); got != test.output {
			assert.Fail(t, "CountSmileyFace(%v) = %v; want %v", test.input, got, test.output)
		}
	}
}
