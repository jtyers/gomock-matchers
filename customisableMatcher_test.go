package gomockmatchers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomisableMatcher(t *testing.T) {
	var tests = []struct {
		name            string
		matcherFunction func(interface{}) bool
		input           interface{}
		expectedResult  bool
	}{
		{
			"should pass for a matcherFunction that returns true",
			func(x interface{}) bool {
				return true
			},
			"my test input",
			true,
		},
		{
			"should fail for a matcherFunction that returns false",
			func(x interface{}) bool {
				return false
			},
			"my test input",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// given
			matcher := CustomisableMatcher(test.matcherFunction)

			// when
			result := matcher.Matches(test.input)

			// then
			require.Equal(t, test.expectedResult, result)
		})
	}
}
