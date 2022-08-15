package utils

type Currency string

func (currency Currency) IsValid() bool {
	switch currency {
	case USD, GBP, NGN:
		return true
	}
	return false
}

const (
	USD Currency = "USD"
	GBP          = "GBP"
	NGN          = "NGN"
)

type Amount struct {
	Value    int64    `json:"value"`
	Currency Currency `json:"currency"`
}
