// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import (
	"fmt"
	"strings"
)

type Function func(args []*Result) *Result

var BuiltinFunctions = make(map[string]Function)

func registerBuiltinFunction(name string, fn Function) {
	BuiltinFunctions[strings.ToUpper(name)] = fn
}

func getBuiltinFunction(name string) Function {
	return BuiltinFunctions[strings.ToUpper(name)]
}

func callBuiltinFunction(name string, args []*Result) *Result {
	fn := getBuiltinFunction(name)
	if fn == nil {
		return newErrorResult(fmt.Errorf("%s is not a function", name))
	}
	return fn(args)
}

func getFunctionArgs(args []*Result) []*Result {
	for i, v := range args {
		a, ok := v.Value.(*AST)
		if ok {
			args[i] = EvalAST(a)
		}
	}
	return args
}
