package utils

var currency = map[string]string{
	"USD": "USD",
	"EUR": "EUR",
	"CAD": "CAD",
}

func IsCurrencyValid(inputCurrency string) bool {
	_, ok := currency[inputCurrency]
	return ok
}
