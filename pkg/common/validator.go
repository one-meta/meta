package common

import (
	"errors"
	"github.com/one-meta/meta/app/ent"
)

// CheckCasbinRule casbin rule type 非 g 或者 p返回错误
func CheckCasbinRule(rule *ent.CasbinRule) error {
	if rule == nil || (rule.Type != "g" && rule.Type != "p") {
		return errors.New("unknown casbin rule type")
	}
	return nil
}
