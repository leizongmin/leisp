// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type OperatorHandler (func(*LangEnv, []Atom) Atom)

type Operator struct {
	symbol      string
	minArgCount int
	maxArgCount int
	handler     OperatorHandler
}
