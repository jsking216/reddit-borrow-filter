package parsers

import (
	"strings"
)

func isRequest(title string) bool {
	return strings.Contains(title, "REQ") && !strings.Contains(strings.ToLower(title), "ARRANGED")
}

func isUSD(title string) bool {
	return strings.Contains(title, "$") || strings.Contains(title, "USD")
}

func ShouldConsider(title string) bool {
	return isRequest(title) && isUSD(title)
}

func ParseRepayDate(title string) {

}
