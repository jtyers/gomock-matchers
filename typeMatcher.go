package gomockmatchers

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

const TypeMatcherDescription = "callback function returns true"

func SameTypeAs(example interface{}) gomock.Matcher {
	return typeMatcher{example}
}

type typeMatcher struct {
	// The function to run to test for matches. Provided by caller.
	example interface{}
}

func (o typeMatcher) Matches(x interface{}) bool {
	return reflect.TypeOf(o.example) == reflect.TypeOf(x)
}

func (o typeMatcher) String() string {
	return TypeMatcherDescription
}
