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
	return fmt.Sprint(r.Value)
}

func (r *Result) Print() {
	if r.Error != nil {
		fmt.Printf("<Error: %s>\n", r.Error)
	} else {
		fmt.Println(r.ToString())
	}
}
