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

func (v *ExpressionValue) IsValue() bool {
	return false
}

func (v *ExpressionValue) ConvertTo(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert expression to %s: does not implement yet", t)
}

func (v *ExpressionValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*ExpressionValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func NewExpressionValue(v *AST) *ExpressionValue {
	return &ExpressionValue{Value: v}
}
