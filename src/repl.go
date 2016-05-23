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
(println {1 2.2 "aa"})
[1 2 3]
`
	parser.Dump(str)

	a := interpreter.Eval(str)
	a.Print()
}
