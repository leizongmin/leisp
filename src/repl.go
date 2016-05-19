// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import (
	"fmt"
	"leisp"
	"strings"
)

func main() {

	reader := strings.NewReader(`
(defun hello {a b}
  ;;(setf c 1/23 'c',)
  (:hello world)
  (format nil (+ a b)))
	`)

	// s := leisp.NewScanner(reader)
	// for {
	// 	tok, lit := s.Scan()
	// 	fmt.Printf("[%d, %d], token=%d, literal='%s'\n", s.Position.Line, s.Position.Column, tok, lit)

	// 	if tok == leisp.TokenIllegal {
	// 		fmt.Printf("Illegal token\n")
	// 		break
	// 	}
	// 	if tok == leisp.TokenEOF {
	// 		fmt.Println("EOF")
	// 		break
	// 	}
	// }

	p := leisp.NewParser(reader)
	if a, e := p.Parse(); e != nil {
		fmt.Println(a, e)
	} else {
		fmt.Println(a)
		if a.Children != nil {
			for i, v := range a.Children {
				fmt.Println(i, *v)
			}
		}
	}
}
