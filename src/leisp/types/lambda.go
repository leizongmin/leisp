// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type LambdaValueInfo struct {
	Names  []string
	Body   []*ExpressionValue
	Source string
}

func NewLambdaValueInfo(names []string, body []*ExpressionValue, source string) *LambdaValueInfo {
	return &LambdaValueInfo{
		Names:  names,
		Body:   body,
		Source: source,
	}
}

type LambdaValue struct {
	Value *LambdaValueInfo
}

func (v *LambdaValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *LambdaValue) GetType() string {
	return "lambda"
}

func (v *LambdaValue) IsValue() bool {
	return false
}

func (v *LambdaValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert keyword to %s: does not implement yet", t)
}

func NewLambdaValue(v *LambdaValueInfo) *LambdaValue {
	return &LambdaValue{Value: v}
}
