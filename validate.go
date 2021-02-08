package postal

import "regexp"

const codeChecker = "^\\d{5}(?:[-\\s]\\d{4})?$"

// IsValid checks whether the supplied request is valid postal code.
func IsValid(code string) bool {
	return regexp.MustCompile(codeChecker).Match([]byte(code))
}
