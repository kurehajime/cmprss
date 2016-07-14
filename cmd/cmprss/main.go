package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kurehajime/cmprss"
)

func main() {
	var text string
	var ok bool
	flag.Parse()
	if len(flag.Args()) == 0 {
		text, ok = readPipe()
	}

	if ok == false {
		text, ok = readFileByArg()
		if ok == false {
			os.Exit(1)
		}
	}

	s := cmprss.Cmprss(text)
	fmt.Println(s)
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
	}
	return "", false
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
