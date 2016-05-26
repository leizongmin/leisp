// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import (
	"fmt"
	"strings"
)

type LambdaValueInfo struct {
	Scope  *Scope
	Names  []string
	Body   []*ExpressionValue
	Source string
}

func NewLambdaValueInfo(scope *Scope, names []string, body []*ExpressionValue, source string) *LambdaValueInfo {
	return &LambdaValueInfo{
		Scope:  scope,
		Names:  names,
		Body:   body,
		Source: source,
	}
}

type LambdaValue struct {
	Value *LambdaValueInfo
}

func (v *LambdaValue) ToString() string {
	return fmt.Sprintf("<Lambda#(%s)#%s>", strings.Join(v.Value.Names, ","), v.Value.Source)
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

func (v *LambdaValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*LambdaValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func NewLambdaValue(v *LambdaValueInfo) *LambdaValue {
	return &LambdaValue{Value: v}
}
