package cssminify

import (
	"regexp"
	"strings"
)

func minifyVal(value string) string {
	value = cleanSpaces(value)

	// Values need special care
	value = cleanHex(value)
	value = cleanUrl(value)
	return value
}

func cleanHex(value string) string {
	re := regexp.MustCompile(`#(\d{6})`)
	matches := re.FindAllString(value, -1)
	for _, hex := range matches {
		if isFull(hex) {
			r := strings.NewReplacer(hex, newHex(hex))
			value = r.Replace(value)
		}
	}
	return value
}

func isFull(hex string) bool {
	cmp := []byte(hex)[1:]
	for _, l := range cmp {
		if cmp[0] != l {
			return false
		}
	}
	return true
}

func newHex(hex string) string {
	return string([]byte(hex)[:4])
}

func cleanUrl(value string) string {
	return value
}