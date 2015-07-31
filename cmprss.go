package main
import(
	"fmt"
	"strings"
	"regexp"
	"os"
	"io/ioutil"
)
func main(){
	var text string
	text, ok := readPipe()
	if ok ==false {
		text, ok = readFileByArg()
		if ok == false {
			os.Exit(1)
		}
	}
	s:=cmprss(text)
	fmt.Println(s)
}
//CnvrtShrtStrng
func cmprss(input string)string{
	re1, _ := regexp.Compile("([bcdfghjklmnpqrstvwxwyzBCDFGHJKLMNPQRSTVWXWYZ])[aiueo]([bcdfghjklmnpqrstvwxwyz])")
	lines:=strings.Split(input,"\n")
	for l:=0;l<len(lines);l++{
		words:= strings.Split(lines[l]," ")
		for i:=0;i<len(words);i++{
			word:=strings.Title(words[i])
			if len(word)>3{
				word = re1.ReplaceAllString(word, "$1$2")
				word = re1.ReplaceAllString(word, "$1$2")
			}
			words[i]=word
		}
		lines[l]= strings.Join(words,"")
	}
	return strings.Join(lines,"\n")
}
//ReadTxtByPpe
func readPipe() (string, bool) {
	stats, _ := os.Stdin.Stat()
	if stats != nil && (stats.Mode()&os.ModeCharDevice) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err.Error())
			return "", false
		}
		return string(bytes), true
	} else {
		return "", false
	}
}
//ReadTxtByFle
func readFileByArg() (string, bool) {
	if len(os.Args) < 2 {
		return "", false
	}
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return "", false
	}
	return string(content), true
}