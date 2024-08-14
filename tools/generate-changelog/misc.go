package main

import (
	"os/exec"
	"path/filepath"
	"slices"
)

var reset = "\033[0m"

// var red = "\033[31m"
// var green = "\033[32m"
var yellow = "\033[33m"

// var blue = "\033[34m"
// var purple = "\033[35m"
// var cyan = "\033[36m"
// var gray = "\033[37m"
// var white = "\033[97m"

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func getModuleRoot() (root string) {
	cmd := exec.Command("go", "env", "GOMOD")
	out, err := cmd.Output()
	die(err)
	return filepath.Dir(string(out)) + "/"
}

// AddLineBreaks is a very naive way to add line breaks
// to a .go file string. It came about as a solution when
// regex replace were too slow to be feasible
func AddLineBreaks(r string) string {
	rr := []rune(r)
	var pa rune
	var insideQuote bool
	var quote rune
	for pos := len(rr) - 1; pos > 0; pos-- {
		char := rr[pos]
		if insideQuote {
			if char == quote {
				if char == '`' {
					insideQuote = false
				} else if pa != '\\' {
					insideQuote = false
				}
			}
		} else {
			if char == '"' || char == '\'' || char == '`' {
				quote = char
				insideQuote = true
			} else {
				if char == ',' || char == '{' {
					rr = slices.Insert(rr, pos+1, '\n')
				}
			}
		}
		pa = char
	}
	return string(rr)
}
