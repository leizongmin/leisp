// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package builtin

import "fmt"

func init() {

	Register("print", Function(func(args []*Result) *Result {
		args = GetArgs(args)
		for _, v := range args {
			fmt.Print(v.Value, " ")
		}
		fmt.Print("\n")
		return newEmptyResult()
	}))

}
