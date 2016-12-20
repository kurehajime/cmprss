package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"io"

	"github.com/kurehajime/cmprss"
)

func main() {
	var r io.Reader
	var ok bool
	flag.Parse()
	if len(flag.Args()) == 0 {
		r, ok = readPipe()
	}

	if ok == false {
		if flag.Arg(0) == "-" {
			text, ok := readStdin()
			r = strings.NewReader(text)
			if ok == false {
				os.Exit(1)
			}
		} else {
			r, ok = readFileByArg()
			if ok == false {
				os.Exit(1)
			}
		}

	}
	cmprss.Cmprss(r, os.Stdout)
}

//ReadTxtByPpe
func readPipe() (io.Reader, bool) {
	stats, _ := os.Stdin.Stat()
	if stats != nil && (stats.Mode()&os.ModeCharDevice) == 0 {
		return os.Stdin, true
	}
	return nil, false
}

//ReadTxtByFle
func readFileByArg() (io.Reader, bool) {
	if len(os.Args) < 2 {
		return nil, false
	}
	r, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	return r, true
}

//ScnInpt
func readStdin() (string, bool) {
	var text string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		text += s.Text() + "\n"
	}
	if s.Err() != nil {
		fmt.Println(s.Err())
		return "", false
	}
	return text, true
}
