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

	str := `(+ 1 2 {"a" 'a' 4 /})`
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
	} else {
		fmt.Println("result", a)
		if a.Children != nil {
			for i, v := range a.Children {
				fmt.Println(i, *v)
			}
		}
	}
}
