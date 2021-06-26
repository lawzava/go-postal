package postal_test

import (
	"testing"

	"github.com/lawzava/go-postal"
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
		res := postal.IsValid(testCase.input)

		assert.Equal(t, testCase.expectedOutput, res, testCase.input)
	}
}
