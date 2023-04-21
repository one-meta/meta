package entity

import (
	"github.com/one-meta/meta/app/ent/extend"
)

type QueryParam struct {
	Search   []string `json:"search"`
	Current  int      `json:"current"`
	PageSize int      `json:"pageSize"`
	Order    string   `json:"order"`
	extend.TimeParam
}

func NewQueryParam() *QueryParam {
	return &QueryParam{Current: 0, PageSize: 20, Order: "-id"}
}
