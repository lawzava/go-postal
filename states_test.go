package postal_test

import (
	"testing"

	"github.com/lawzava/go-postal"
	"github.com/stretchr/testify/assert"
)

func TestFindState(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input          string
		expectedOutput postal.State
	}{
		{"59000", postal.Montana},
		{"76512", postal.Texas},
		{"", ""},
		{"00033", ""},
	}

	for _, testCase := range testCases {
		res, _ := postal.FindState(testCase.input)

		assert.Equal(t, testCase.expectedOutput, res, testCase.input)
	}
}
