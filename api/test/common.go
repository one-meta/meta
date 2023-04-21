package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/one-meta/meta/pkg/common"
)

type Expected struct {
	StatusCode   int           `json:"statusCode"`
	ResponseData string        `json:"responseData"`
	Result       common.Result `json:"result"`
}
type Assert struct {
	EqualStatusCode      bool `json:"equalStatusCode"`
	EqualResponseData    bool `json:"equalResponseData"`
	ContainsResponseData bool `json:"containsResponseData"`
	EqualFlag            bool `json:"equalFlag"`
	EqualCode            bool `json:"equalCode"`
	EqualMessage         bool `json:"equalMessage"`
	ContainsMessage      bool `json:"containsMessage"`
	ContainsData         bool `json:"ContainsData"`
	PageGrater           bool `json:"pageGrater"`
}
type CaseRule struct {
	Api        string            `json:"api"`
	HttpMethod string            `json:"httpMethod"`
	Header     map[string]string `json:"header"`
	Action     string            `json:"action"`
	UrlData    string            `json:"urlData"`
	BodyData   any               `json:"bodyData"`
	Expected   *Expected         `json:"expected"`
	Assert     *Assert           `json:"assert"`
}

var (
	notFoundResult       = &common.Result{Success: false, Message: "Not Found"}
	errorFoundResult     = &common.Result{Success: false, Message: "not found"}
	successResult        = &common.Result{Success: true}
	accessDenyResult     = &common.Result{Success: false, Message: "Access denied"}
	permissionDenyResult = &common.Result{Success: false, Message: "Permission denied"}

	notFoundExpectedResult       = &Expected{StatusCode: 404, Result: *notFoundResult}
	errorFoundExpectedResult     = &Expected{StatusCode: 200, Result: *errorFoundResult}
	successExpectedResult        = &Expected{StatusCode: 200, Result: *successResult}
	accessDenyExpectedResult     = &Expected{StatusCode: 403, Result: *accessDenyResult}
	permissionDenyExpectedResult = &Expected{StatusCode: 401, Result: *permissionDenyResult}

	assertEqualContains       = &Assert{EqualFlag: true, EqualCode: true, EqualStatusCode: true, ContainsMessage: true, ContainsData: true}
	assertEqualContainsGrater = &Assert{EqualFlag: true, EqualCode: true, EqualStatusCode: true, ContainsMessage: true, ContainsData: true, PageGrater: true}
	assertEqual               = &Assert{EqualFlag: true, EqualCode: true, EqualStatusCode: true}

	baseApi  = "/api/v1"
	fiberApp *fiber.App
	f        func()
)
