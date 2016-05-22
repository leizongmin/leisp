// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import "leisp/parser"

func main() {

	str := `
{1 2 3}
(defun print [msg] (format "Hello, " msg))
`
	parser.Dump(str)

	// fmt.Println()
	// r := leisp.Eval(str)
	// r.Print()
}
