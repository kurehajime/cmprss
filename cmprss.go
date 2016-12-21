// Package cmprss : TheSntnceToLssyCmprssion.
package cmprss

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

//Cmprss : CnvrtShrtStrng
func Cmprss(r io.Reader, w io.Writer) {
	re, _ := regexp.Compile("([bcdfghjklmnpqrstvwxwyzBCDFGHJKLMNPQRSTVWXWYZ])[aiueo]([bcdfghjklmnpqrstvwxwyz])")
	r2 := bufio.NewReader(r)
	for {
		// line includes '\n'.
		text, err := r2.ReadString('\n')
		words := strings.Split(text, " ")
		for i := 0; i < len(words); i++ {
			word := strings.Title(words[i])
			if len(word) > 3 {
				word = re.ReplaceAllString(re.ReplaceAllString(word, "$1$2"), "$1$2")
			}
			words[i] = word
		}
		w.Write([]byte(strings.Join(words, "")))
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
	return
}
