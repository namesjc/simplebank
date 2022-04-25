package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/namesjc/simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if Currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check currency is supported
		return util.IsSupportCurrency(Currency)
	}
	return false
}
