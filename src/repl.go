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

	reader := strings.NewReader("hello")

	s := leisp.NewScanner(reader)

	fmt.Println(s.Scan())
}
