package gomockmatchers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type SomeStruct struct {
}

type SomeStruct2 struct {
}

type SomeInterface interface {
}

func TestTypeMatcher(t *testing.T) {
	var tests = []struct {
		name           string
		example        interface{}
		input          interface{}
		expectedResult bool
	}{
		{
			"should pass for matching simple types",
			"a string",
			"my test input",
			true,
		},
		{
			"should fail for different simple types",
			123,
			"my test input",
			false,
		},
		{
			"should pass for matching struct types",
			SomeStruct{},
			SomeStruct{},
			true,
		},
		{
			"should fail for different struct types",
			SomeStruct{},
			SomeStruct2{},
			false,
		},
		{
			"should pass for matching *struct types",
			&SomeStruct{},
			&SomeStruct{},
			true,
		},
		{
			"should fail for different *struct types",
			&SomeStruct{},
			&SomeStruct2{},
			false,
		},
		{
			"should pass for matching interface types",
			context.Background(),
			context.TODO(),
			true,
		},
		{
			"should pass for matching interface types via new()",
			new(context.Context),
			context.TODO(),
			true,
		},
		{
			"should fail for different interface types",
			context.Background(),
			"", // implements Stringer
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// given
			matcher := SameTypeAs(test.example)

			// when
			result := matcher.Matches(test.input)

			// then
			require.Equal(t, test.expectedResult, result)
		})
	}
}
