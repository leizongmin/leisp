// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import "fmt"

type Result struct {
	Error error
	Value interface{}
}

func newResult(value interface{}) *Result {
	return &Result{
		Value: value,
	}
}

func newEmptyResult() *Result {
	return newResult(nil)
}

func newErrorResult(err error) *Result {
	return &Result{
		Error: err,
	}
}

func (r *Result) ToString() string {
	if r.Error != nil {
		return fmt.Sprintf("<Error#\"%s\">", r.Error)
	}
	if arr, ok := r.Value.([]*AST); ok {
		arr2 := make([]interface{}, len(arr))
		for i, v := range arr {
			arr2[i] = v.Value
		}
		r.Value = arr2
	}
	return fmt.Sprint(r.Value)
}

func (r *Result) Print() {
	fmt.Println(r.ToString())
}
