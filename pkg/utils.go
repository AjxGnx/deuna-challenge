package pkg

import (
	"strings"
)

func GetPaymentMethod(paymentMethod string) string {
	if strings.Contains(paymentMethod, "card") {
		return "card"
	}

	return paymentMethod
}
