package gomockmatchers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func stringPtr(s string) *string {
	return &s
}

type FamilyStruct struct {
	LastName string
	ID       string

	Ptr *string
}

type FamilyMemberStruct struct {
	FamilyStruct // embedded

	FirstName string
}

func TestPartialStructMatcher(t *testing.T) {
	ptr := stringPtr("foo")

	var tests = []struct {
		name           string
		matcherStruct  interface{}
		input          interface{}
		expectedResult bool
	}{
		{
			"should pass for equal structs",
			FamilyStruct{LastName: "bar", ID: "01"},
			FamilyStruct{LastName: "bar", ID: "01"},
			true,
		},
		{
			"should fail for unequal structs",
			FamilyStruct{LastName: "foo", ID: "01"},
			FamilyStruct{LastName: "bar", ID: "01"},
			false,
		},
		{
			"should pass for equal fields in a struct with more fields set",
			FamilyStruct{LastName: "bar"},
			FamilyStruct{LastName: "bar", ID: "01"},
			true,
		},
		{
			"should fail for nonequal fields in a struct with more fields set",
			FamilyStruct{LastName: "bar"},
			FamilyStruct{LastName: "foo", ID: "01"},
			false,
		},
		{
			"should pass for equal fields across parent/embedded struct",
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}, FirstName: "baz"},
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}, FirstName: "baz"},
			true,
		},
		{
			"should fail for unequal fields in embedded struct",
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}, FirstName: "bar"},
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}, FirstName: "baz"},
			false,
		},
		{
			"should fail for unequal fields in parent struct",
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "foo"}, FirstName: "baz"},
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}, FirstName: "baz"},
			false,
		},
		{
			"should pass for equal fields in embedded struct",
			FamilyMemberStruct{FirstName: "baz"},
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}, FirstName: "baz"},
			true,
		},
		{
			"should pass for equal fields in parent struct",
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}},
			FamilyMemberStruct{FamilyStruct: FamilyStruct{LastName: "bar"}, FirstName: "baz"},
			true,
		},
		{
			"should pass for exactly equal pointers",
			FamilyMemberStruct{FamilyStruct: FamilyStruct{Ptr: ptr}},
			FamilyMemberStruct{FamilyStruct: FamilyStruct{Ptr: ptr}},
			true,
		},
		{
			"should pass for pointers to equal values",
			FamilyMemberStruct{FamilyStruct: FamilyStruct{Ptr: stringPtr("foo")}},
			FamilyMemberStruct{FamilyStruct: FamilyStruct{Ptr: stringPtr("foo")}},
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// given
			matcher := PartialStructMatcher(test.matcherStruct)

			// when
			result := matcher.Matches(test.input)

			// then
			require.Equal(t, test.expectedResult, result)
		})
	}
}
