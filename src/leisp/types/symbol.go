// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type SymbolValue struct {
	Value string
}

func (t *SymbolValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *SymbolValue) GetType() string {
	return "symbol"
}

func NewSymbol(v string) *SymbolValue {
	return &SymbolValue{Value: v}
}
