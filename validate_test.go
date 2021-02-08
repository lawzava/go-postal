package postal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input          string
		expectedOutput bool
	}{
		{"12349", true},
		{"01519", true},
		{"15134 2354", true},
		{"12", false},
		{"5615", false},
		{"015", false},
		{"", false},
		{"dsafnan", false},
	}

	for _, testCase := range testCases {
		res := IsValid(testCase.input)

		assert.Equal(t, testCase.expectedOutput, res, testCase.input)
	}
}
