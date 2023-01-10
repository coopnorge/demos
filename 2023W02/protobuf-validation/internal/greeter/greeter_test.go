package greeter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeFunction(t *testing.T) {
	type ResultType struct {
		greeting string
		err      error
	}

	type TestCase struct {
		name           string
		inputName      string
		expectedResult string
		errorRegex     string
		validate       func(t *testing.T, testCase *TestCase, result ResultType)
	}
	type validateFunc func(t *testing.T, testCase *TestCase, result ResultType)

	var checkTest validateFunc = func(t *testing.T, testCase *TestCase, result ResultType) {
		if len(testCase.errorRegex) > 0 {
			assert.Error(t, result.err)
			assert.Regexp(t, testCase.errorRegex, result.err.Error())
		} else {
			if result.err != nil {
				t.Fatal(result.err)
			}
			assert.Equal(t, result.greeting, testCase.expectedResult)
			assert.NoError(t, result.err)
		}
	}

	testCases := []TestCase{
		{
			name:           "everything okay",
			inputName:      "Bob",
			expectedResult: "Hi Bob!",
		},
		{
			name:           "everything okay",
			inputName:      "John",
			expectedResult: "",
			errorRegex:     "your name can't be John",
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(fmt.Sprintf("%v", testCase.name), func(t *testing.T) {
			greeting, err := DoSomething(testCase.inputName)

			if testCase.validate != nil {
				testCase.validate(t, &testCase, ResultType{greeting, err})
			} else {
				checkTest(t, &testCase, ResultType{greeting, err})
			}
		})
	}
}
