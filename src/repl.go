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
(def aa "hello, world")
(println aa 123 "ok" :haha def)
`
	parser.Dump(str)

	a := interpreter.Eval(str)
	a.Print()
}
