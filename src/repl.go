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
(type-of 1/2)
`
	parser.Dump(str)

	a := interpreter.Eval(str)
	a.Print()
}
