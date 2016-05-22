// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package parser

import (
	"fmt"
	"strings"
)

func Dump(s string) {
	fmt.Println(s)
	fmt.Println("----------------------------------------")
	p := NewParser(strings.NewReader(s))
	for {
		if a, e := p.Parse(); e != nil {
			pos := p.GetPosition()
			fmt.Printf("Error: %s at line %d,%d\n", e.Error(), pos.Line, pos.Column)
			break
		} else if a.IsEOF() {
			fmt.Println("EOF")
			break
		} else {
			a.Dump()
			fmt.Println("----------------------------------------")
		}
	}
}
