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

	str := `(+ 1 2.2e5 {"a" 'a' 4 / [1] []
			  ;;这里使用keyword
			  (:name my-name)})
			(setf a 1.1)`
	reader := strings.NewReader(str)

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
		fmt.Println("result", a, e)
		pos := p.GetPosition()
		fmt.Printf("Error: %s at line %d,%d\n", e.Error(), pos.Line, pos.Column)
	} else {
		fmt.Println("result", a)
		a.Dump()
	}
}
