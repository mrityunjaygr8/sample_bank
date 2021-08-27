package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mrityunjaygr8/sample_bank/utils"
)

var currencyValid validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return utils.IsCurrencyValid(currency)
	}
	return false
}
