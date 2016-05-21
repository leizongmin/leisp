// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package builtin

import (
	"fmt"
	"strings"
)

type Function func(args []*Result) *Result

var Functions = make(map[string]Function)

func Register(name string, fn Function) {
	Functions[strings.ToUpper(name)] = fn
}

func Get(name string) Function {
	return Functions[strings.ToUpper(name)]
}

func Call(name string, args []*Result) *Result {
	fn := Get(name)
	if fn == nil {
		return newErrorResult(fmt.Errorf("%s is not a function", name))
	}
	return fn(args)
}

func GetArgs(args []*Result) []*Result {
	for i, v := range args {
		a, ok := v.Value.(*AST)
		if ok {
			args[i] = EvalAST(a)
		}
	}
	return args
}
