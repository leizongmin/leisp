// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type ExpressionValue struct {
	Value *AST
}

func (t *ExpressionValue) ToString() string {
	return t.Value.ToString()
}

func (t *ExpressionValue) GetType() string {
	return "integer"
}

func NewExpression(v *AST) *ExpressionValue {
	return &ExpressionValue{Value: v}
}
