package common

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/one-meta/meta/app/entity"
)

// QueryParser 将url中参数解析成实体和queryParam
func QueryParser(c *fiber.Ctx, pojo any) (*entity.QueryParam, error) {
	err := c.QueryParser(pojo)
	if err != nil {
		return nil, err
	}
	qp := entity.NewQueryParam()
	err = c.QueryParser(qp)
	if err != nil {
		return nil, err
	}
	//排序字段转换
	//createdAt => created_at
	qp.Order = Snake(qp.Order)

	//前端分页转换
	if qp.Current <= 1 {
		qp.Current = 0
	} else {
		qp.Current = (qp.Current - 1) * qp.PageSize
	}
	return qp, nil
}

// QueryParamParser 将url中参数解析成queryParam
func QueryParamParser(c *fiber.Ctx, queryParam *entity.QueryParam) error {
	err := c.QueryParser(queryParam)
	if err != nil {
		return err
	}
	return nil
}

// QueryStructParser 将url中参数解析成实体
func QueryStructParser(c *fiber.Ctx, entity any) error {
	err := c.QueryParser(entity)
	if err != nil {
		return err
	}
	return nil
}

// BodyParser 将body中数据解析成实体
func BodyParser(c *fiber.Ctx, entity any) error {
	err := c.BodyParser(entity)
	if err != nil {
		return err
	}
	return nil
}

// RequestBodyParser 将request body的json数据解析成实体
func RequestBodyParser(c *fiber.Ctx, entity any) error {
	body := c.Request().Body()
	err := sonic.Unmarshal(body, entity)
	if err != nil {
		return err
	}
	return nil
}
