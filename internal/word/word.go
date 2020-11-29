package word

import (
	"strings"
	"unicode"
)

//全轉大寫
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//全轉小寫
func ToLower(s string) string {
	return strings.ToLower(s)
}

//底線轉大寫駝峰單字
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//底線轉小寫駝峰單字
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//駝峰單字轉底線單字
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
