// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"leisp/interpreter"
	"os"
	"strings"

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

func printUsage() {

	fmt.Printf("%s\n", `
Usage: leisp [options] [<file1> <file2> ...]

Options:
	`)

	flag.PrintDefaults()
	os.Exit(0)

}

func main() {

	flag.Usage = printUsage
	flag.Parse()

	printWelcome()

	files := flag.Args()
	for _, f := range files {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		r := interpreter.Eval(nil, string(content))
		if r.IsError() {
			fmt.Println(r.ToString())
			os.Exit(2)
		}
	}

	startREPL()
}

func createLiner() *liner.State {

	rl := liner.NewLiner()
	rl.SetCtrlCAborts(true)
	rl.SetMultiLineMode(true)

	rl.SetTabCompletionStyle(liner.TabCircular)
	rl.SetWordCompleter(func(line string, pos int) (head string, completions []string, tail string) {
		if pos == 0 {
			head, tail = "(", ")"
			completions = []string{""}
		} else {
			// fmt.Println("")
			// str := line[pos:]
			// fmt.Println(str)
			// for _, n := range interpreter.Scope.Keys() {
			// 	fmt.Print(n, "\t")
			// 	if strings.Index(n, str) == 0 {
			// 		completions = append(completions, n)
			// 	}
			// }
			// fmt.Println("")
		}
		return head, completions, tail
	})

	return rl
}

func startREPL() {

	rl := createLiner()

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

		buffer = strings.TrimSpace(buffer + line + " ")
		if len(brackets) == 0 && !isString {

			a := interpreter.Eval(nil, buffer)
			a.Print()

			rl.AppendHistory(buffer)
			buffer = ""
		}

	}

	rl.Close()
}
