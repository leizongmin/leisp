// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import (
	"leisp/interpreter"
	"leisp/parser"
)

func main() {

	str := `
{1 2 3}
(defun print [msg] (format "Hello, " msg))

1 "xyz" 'a' nil t
`
	parser.Dump(str)

	a := interpreter.Eval(str)
	// r.Print()
}
