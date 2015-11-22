package testilla

import (
	"fmt"
	"net/url"
)

// IApiTest defines an interface of API test definition.
//
// Responsibility of api test is to provide some meta data about API endpoint and
// collection of test cases that describe concrete usage of that endpoint.
type IApiTest interface {
	Method() string
	Description() string
	Path() string
	TestCases() []ApiTestCase
}

// ApiTestCase provides use case of some API endpoint: input and expected output.
//
// Test case knows nothing about API endpoint itself. Path of an API endpoint and
// its description must be provided by ApiTest.
//
// Ideally each different error that API endpoint can return should
// be described by a test case.
type ApiTestCase struct {
	Description string

	Headers     map[string]ApiTestCaseParam
	QueryParams map[string]ApiTestCaseParam
	RequestBody interface{}

	ExpectedHttpCode int
	ExpectedHeaders  map[string]string
	ExpectedData     interface{}
}

// ApiTestCaseParam defines a parameter that is used in headers or URL query
// of the API request
type ApiTestCaseParam struct {
	Value    interface{}
	Required bool
}

// Url generates full URL to API endpoint for given test case.
// urlpath must provide full URL to the endpoint with no query parameters.
func (testCase *ApiTestCase) Url(urlpath string) (string, error) {
	u, err := url.Parse(urlpath)
	if err != nil {
		return "", err
	}
	if testCase.QueryParams != nil {
		query := url.Values{}
		for key, param := range testCase.QueryParams {
			valueStr := fmt.Sprintf("%v", param.Value)
			query.Set(key, valueStr)
		}
		u.RawQuery = query.Encode()
	}

	return u.String(), nil
}