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
  (format nil (+ a b)))
	`)

	s := leisp.NewScanner(reader)

	for {
		tok, lit := s.Scan()
		fmt.Printf("token=%d, literal='%s'\n", tok, lit)

		if tok == leisp.TokenIllegal {
			fmt.Printf("Illegal token\n")
			break
		}
		if tok == leisp.TokenEOF {
			fmt.Println("EOF")
			break
		}
	}
}
