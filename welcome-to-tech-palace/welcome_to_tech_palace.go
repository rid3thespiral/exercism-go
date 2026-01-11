package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var result string
	for i := 0; i < numStarsPerLine; i++ {
		result += "*"
	}
	return result + "\n" + welcomeMsg + "\n" + result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	res := strings.Replace(oldMsg, "*", "", -1)
	res2 := strings.Replace(res, "\n", "", -1)
	return strings.TrimSpace(res2)
}
