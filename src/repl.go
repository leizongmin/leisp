// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import (
	"fmt"
	"leisp/parser"
	"strings"
)

func main() {

	str := `
{1 2 3}
`
	fmt.Println(str)

	reader := strings.NewReader(str)
	p := parser.NewParser(reader)
	if a, e := p.Parse(); e != nil {
		pos := p.GetPosition()
		fmt.Printf("Error: %s at line %d,%d\n", e.Error(), pos.Line, pos.Column)
	} else {
		a.Dump()
	}

	// fmt.Println()
	// r := leisp.Eval(str)
	// r.Print()
}
