package cmprss

import (
	"regexp"
	"strings"
)

//Cmprss : CnvrtShrtStrng
func Cmprss(input string) string {
	re1, _ := regexp.Compile("([bcdfghjklmnpqrstvwxwyzBCDFGHJKLMNPQRSTVWXWYZ])[aiueo]([bcdfghjklmnpqrstvwxwyz])")
	lines := strings.Split(input, "\n")
	for l := 0; l < len(lines); l++ {
		words := strings.Split(lines[l], " ")
		for i := 0; i < len(words); i++ {
			word := strings.Title(words[i])
			if len(word) > 3 {
				word = re1.ReplaceAllString(word, "$1$2")
				word = re1.ReplaceAllString(word, "$1$2")
			}
			words[i] = word
		}
		lines[l] = strings.Join(words, "")
	}
	return strings.Join(lines, "\n")
}
