// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package builtin

func init() {

	Register("setf", Function(func(args []*Result) *Result {
		return newResult("setf")
	}))

}
