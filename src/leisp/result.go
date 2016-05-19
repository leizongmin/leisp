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

func newErrorResult(err error) *Result {
	return &Result{
		Error: err,
	}
}

func (r *Result) ToString() string {
	return fmt.Sprint(r.Value)
}

func (r *Result) Print() {
	fmt.Println(r.ToString())
}
