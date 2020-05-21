package gomockmatchers

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
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
	exampleType := reflect.TypeOf(o.example)
	xType := reflect.TypeOf(x)

	var exampleElemType reflect.Type
	if exampleType.Kind() == reflect.Ptr {
		exampleElemType = exampleType.Elem()
	}

	if exampleType.Kind() == reflect.Interface || (exampleElemType != nil && exampleElemType.Kind() == reflect.Interface) {
		//exampleType.Elem().Kind()
		// for interface examples, check AssignableTo
		fmt.Printf("exampleType %v\n", exampleType)
		fmt.Printf("xType %v\n", xType)
		return xType.AssignableTo(exampleType) || xType.AssignableTo(exampleElemType)

	} else {
		// for concrete types, check equality
		return exampleType == xType
	}
}

func (o typeMatcher) String() string {
	return TypeMatcherDescription
}
