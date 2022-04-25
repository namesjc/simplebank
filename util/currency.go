package util

const (
	USD  = "USD"
	EUR  = "EUR"
	YUAN = "YUAN"
)

// IsSupportCurrency returns true if the currency is supported
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, YUAN:
		return true
	}
	return false
}
