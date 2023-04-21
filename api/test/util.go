package test

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"github.com/one-meta/meta/app/entity/config"
	"github.com/one-meta/meta/pkg/common"
	"github.com/one-meta/meta/pkg/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func runTest(t *testing.T, caseRule *CaseRule) any {
	req := genCaseRuleRequest(caseRule)
	resp, result, resultData := fiberTest(t, fiberApp, req)
	//fmt.Println("resp ", resp)
	//fmt.Println("result ", result)
	//fmt.Println("resultData ", resultData)
	assertData(t, caseRule, resp, result, resultData)
	return result.Data
}

func assertData(t *testing.T, c *CaseRule, resp *http.Response, result *common.Result, containStr string) {
	if c.Assert.EqualCode {
		assert.Equal(t, c.Expected.StatusCode, resp.StatusCode, "响应状态码与期望相同")
	}
	if c.Assert.EqualResponseData {
		assert.Equal(t, c.Expected.ResponseData, resp.Body, "响应数据与期望相同")
	}
	if c.Assert.ContainsResponseData {
		assert.Contains(t, resp.Body, c.Expected.ResponseData, "响应数据中包含期望")
	}
	if c.Assert.EqualFlag {
		assert.Equal(t, c.Expected.Result.Success, result.Success, "返回结果的flag与期望相同")
	}
	if c.Assert.EqualStatusCode {
		assert.Equal(t, c.Expected.Result.Code, result.Code, "返回结果的状态码与期望相同")
	}
	if c.Assert.EqualMessage {
		assert.Equal(t, c.Expected.Result.Message, result.Message, "返回结果的message与期望相同")
	}
	if c.Assert.ContainsMessage {
		assert.Contains(t, result.Message, c.Expected.Result.Message, "返回结果的message包含期望的msg")
	}
	if c.Assert.ContainsData {
		assert.Contains(t, containStr, c.Expected.ResponseData, "返回的结果Data数据中包含期望")
	}
	if c.Assert.PageGrater {
		assert.GreaterOrEqual(t, result.Total, 5, "返回的分页数据大于期望")
	}
}

func fiberTest(t *testing.T, fiberApp *fiber.App, req *http.Request) (*http.Response, *common.Result, string) {
	resp, err := fiberApp.Test(req, -1)
	if err != nil {
		t.Error(err)
	}
	result := printResult(resp)
	resultData, err := sonic.Marshal(result.Data)
	if err != nil {
		log.Fatal(err)
	}
	return resp, result, string(resultData)
}

func genCaseRuleRequest(c *CaseRule) *http.Request {
	bodyData := c.BodyData
	if c.Api != "" {
		var body io.Reader
		if bodyData != nil {
			// fmt.Println("body data", bodyData)
			marshal, err := sonic.Marshal(bodyData)
			if err != nil {
				log.Fatal(err)
			}
			body = strings.NewReader(string(marshal))
		}
		var api string
		if c.UrlData != "" {
			//?拼接的参数
			if strings.Contains(c.UrlData, "=") {
				api = c.Api + "?" + c.UrlData
			} else {
				//restful 参数
				api = c.Api + "/" + c.UrlData
			}
		} else {
			api = c.Api
		}
		req := httptest.NewRequest(c.HttpMethod, api, body)
		req.Header.Add("Content-Type", "application/json")
		if len(c.Header) != 0 {
			for k, v := range c.Header {
				req.Header.Add(k, v)
			}
		}
		return req
	}
	return nil
}

func printResult(resp *http.Response) *common.Result {
	bytes, _ := io.ReadAll(resp.Body)
	result := &common.Result{}
	err := sonic.Unmarshal(bytes, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func initFiber(enableLogin bool) (*fiber.App, func()) {
	err := util.LoadConfig("../../resource")
	if err != nil {
		log.Fatal(err)
	}
	config.CFG.Auth.Enable = enableLogin
	config.CFG.Auth.Casbin.ModelPath = "../../resource"
	config.CFG.Ent.DebugMode = false
	fiberApp, router, _, f := util.InjectApp()
	util.RegisterRouter(router, fiberApp)
	return fiberApp, f
}
func getDataMapId(data any) int {
	dataMap := data.(map[string]any)
	return int(dataMap["id"].(float64))
}
