package common

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type StatusCode int

// Result 响应结果
type Result struct {
	Data    any        `json:"data,omitempty"`
	Code    StatusCode `json:"code,omitempty"`
	Total   int        `json:"total,omitempty"`
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	//Current  int        `json:"current,omitempty"`
	//PageSize int        `json:"pageSize,omitempty"`
}

type DeleteItem struct {
	Ids []int `json:"ids,omitempty"`
}
type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func NewResult(c *fiber.Ctx, err error, data ...any) error {
	if err != nil {
		if err.Error() == "login error" {
			return errMsg(c, errors.New("mismatch username or password"))
		}
		return errMsg(c, err)
	} else {
		r := &Result{Success: true}
		if data != nil {
			if len(data) == 1 {
				r.Data = data[0]
			} else {
				r.Data = data
			}
		}
		return parser(c, r)
	}
}

func NewErrorWithStatusCode(c *fiber.Ctx, err string, statusCode int) error {
	result := &Result{
		Success: false,
		Message: err,
	}
	c.Status(statusCode)
	return parser(c, result)
}

func NewPageResult(c *fiber.Ctx, err error, count int, result any) error {
	if err != nil {
		return errMsg(c, err)
	} else {
		return parser(c, &Result{
			Success: true,
			Data:    result,
			Total:   count,
		})

	}
}

func errMsg(c *fiber.Ctx, err error) error {
	r := &Result{Success: false, Message: err.Error()}
	if err == nil {
		r.Message = "error"
	}
	return parser(c, r)
}

// parser
func parser(c *fiber.Ctx, result *Result) error {
	return c.JSON(result)
}
