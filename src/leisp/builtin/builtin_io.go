// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import "fmt"

func init() {

	registerBuiltinFunction("print", Function(func(args []*Result) *Result {
		args = getFunctionArgs(args)
		for _, v := range args {
			fmt.Print(v.Value, " ")
		}
		fmt.Print("\n")
		return newEmptyResult()
	}))

}
