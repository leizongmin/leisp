// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import (
	"fmt"
	"leisp/interpreter"
	"leisp/parser"

	"github.com/peterh/liner"
)

func main() {

	str := `
(defvar aa 456)
(println aa 123 "ok" :haha 'defvar)
(println {1 2.2 "aa"})
(println [1 2 3])
(println '(list 1 2 3))
(println (str 1 "2" 3.3 :4) (/ 1 2 34) (^ 2 10 2))
(defn add [a b]
  (println "arguments:" a b)
  (+ a b))
(println (add 123 aa))
(println this)
`
	parser.Dump(str)

	a := interpreter.Eval(nil, str)
	a.Print()

	rl := liner.NewLiner()

	for {

		line, err := rl.Prompt("leisp> ")
		if err != nil {
			fmt.Println("Error: %s", err)
			break
		}

		if line == "" {
			continue
		}

		a := interpreter.Eval(nil, line)
		a.Print()
	}

}
