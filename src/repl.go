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
(defvar aa 456)
;(println aa 123 "ok" :haha def)
;(println {1 2.2 "aa"})
;(println [1 2 3])
;(println '(list 1 2 3))
(println (str 1 "2" 3.3 :4) (/ 1 2 34) (^ 2 10 2))
(defn add [a b]
  (println "arguments:" a b)
  (+ a b))
(println (add 123 aa))
(println this)
`
	parser.Dump(str)

	a := interpreter.Eval(str)
	a.Print()
}
