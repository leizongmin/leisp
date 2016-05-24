// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import (
	"fmt"
	"leisp/interpreter"

	"github.com/peterh/liner"
)

func makeIndentString(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "  "
	}
	return s
}

func printWelcome() {
	fmt.Println(`
##       ######## ####  ######  ########
##       ##        ##  ##    ## ##     ##
##       ##        ##  ##       ##     ##
##       ######    ##   ######  ########
##       ##        ##        ## ##
##       ##        ##  ##    ## ##
######## ######## ####  ######  ##

Welcome to leisp 0.0.1

Copyright (c) 2016 Zongmin Lei <http://ucdok.com>

Type (help) and hit Enter for context help.
Press Ctrl+C to Exit.
	`)
}

func main() {

	printWelcome()

	// 	str := `
	// (defvar aa 456)
	// (println aa 123 "ok" :haha 'defvar)
	// (println {1 2.2 "aa"})
	// (println [1 2 3])
	// (println '(list 1 2 3))
	// (println (str 1 "2" 3.3 :4) (/ 1 2 34) (^ 2 10 2))
	// (defn add [a b]
	//   (println "arguments:" a b)
	//   (+ a b))
	// (println (add 123 aa))
	// (println this)
	// `

	// 	parser.Dump(str)

	// 	a := interpreter.Eval(nil, str)
	// 	a.Print()

	rl := liner.NewLiner()
	rl.SetCtrlCAborts(true)

	brackets := make([]rune, 0)
	isString := false
	buffer := ""

	for {

		prompt := "leisp> "
		if len(brackets) > 0 {
			prompt = "       " + makeIndentString(len(brackets))
		}

		line, err := rl.Prompt(prompt)
		if err != nil {
			if err.Error() == "prompt aborted" {
				break
			}
			fmt.Println(err)
			break
		}

		if line == "" {
			continue
		}

		isEscape := false
		for _, r2 := range line {
			if isString {
				if isEscape {
					isEscape = false
				} else {
					if r2 == '\\' {
						isEscape = true
					} else if r2 == '"' {
						isString = false
					}
				}
			} else {
				if r2 == '{' || r2 == '[' || r2 == '(' {
					brackets = append(brackets, r2)
				} else if r2 == '}' || r2 == ']' || r2 == ')' {
					r := brackets[len(brackets)-1]
					if r == '{' && r2 == '}' {
						brackets = brackets[:len(brackets)-1]
					} else if r == '[' && r2 == ']' {
						brackets = brackets[:len(brackets)-1]
					} else if r == '(' && r2 == ')' {
						brackets = brackets[:len(brackets)-1]
					} else {
						fmt.Printf("Error: %s does not matched %s\n", string(r), string(r2))
					}
				} else if r2 == '"' {
					isString = true
				}
			}
		}

		buffer += " " + line
		if len(brackets) == 0 && !isString {
			a := interpreter.Eval(nil, buffer)
			a.Print()
			buffer = ""
		}

	}

	rl.Close()

}
