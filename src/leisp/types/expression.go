// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type ExpressionValue struct {
	Value *AST
}

func (v *ExpressionValue) ToString() string {
	return v.Value.ToString()
}

func (v *ExpressionValue) GetType() string {
	return "integer"
}

func (v *ExpressionValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert expression to %s: does not implement yet", t)
}

func NewExpression(v *AST) *ExpressionValue {
	return &ExpressionValue{Value: v}
}
