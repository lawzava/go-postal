package postal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindState(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input          string
		expectedOutput State
	}{
		{"59000", Montana},
		{"76512", Texas},
		{"", ""},
		{"00033", ""},
	}

	for _, testCase := range testCases {
		res, _ := FindState(testCase.input)

		assert.Equal(t, testCase.expectedOutput, res, testCase.input)
	}
}
