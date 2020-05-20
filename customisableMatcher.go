package gomockmatchers

import (
	"github.com/golang/mock/gomock"
)

const CustomisableMatcherMatcherDescription = "callback function returns true"

// https://github.com/golang/mock/issues/351#issuecomment-560177910
func CustomisableMatcher(matcherFunction func(arg interface{}) bool) gomock.Matcher {
	return customisableMatcher{matcherFunction: matcherFunction}
}

type customisableMatcher struct {
	// The function to run to test for matches. Provided by caller.
	matcherFunction func(arg interface{}) bool
}

func (o customisableMatcher) Matches(x interface{}) bool {
	return o.matcherFunction(x)
}

func (o customisableMatcher) String() string {
	return CustomisableMatcherMatcherDescription
}
