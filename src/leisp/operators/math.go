// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package builtin

import "fmt"

func init() {

	Register("+", Function(func(args []*Result) *Result {

		if len(args) < 1 {
			return newResult(0)
		}

		args = GetArgs(args)
		var ri int64
		var rf float64
		isInt := true

		if vi, ok := args[0].Value.(int64); ok {
			ri = vi
		} else if vf, ok := args[0].Value.(float64); ok {
			isInt = false
			rf = vf
		}

		for _, v := range args[1:] {
			if vi, ok := v.Value.(int64); ok {
				if isInt {
					ri += vi
				} else {
					rf += float64(vi)
				}
				continue
			} else {
				rf = float64(ri)
				isInt = false
			}
			if vf, ok := v.Value.(float64); ok {
				rf += vf
			}
		}

		if isInt {
			return newResult(ri)
		}
		return newResult(rf)
	}))

	Register("-", Function(func(args []*Result) *Result {

		if len(args) < 1 {
			return newResult(0)
		}

		args = GetArgs(args)
		var ri int64
		var rf float64
		isInt := true

		if vi, ok := args[0].Value.(int64); ok {
			ri = vi
		} else if vf, ok := args[0].Value.(float64); ok {
			isInt = false
			rf = vf
		}

		for _, v := range args[1:] {
			if vi, ok := v.Value.(int64); ok {
				if isInt {
					ri -= vi
				} else {
					rf -= float64(vi)
				}
				continue
			} else {
				rf = float64(ri)
				isInt = false
			}
			if vf, ok := v.Value.(float64); ok {
				rf -= vf
			}
		}

		if isInt {
			return newResult(ri)
		}
		return newResult(rf)
	}))

	Register("*", Function(func(args []*Result) *Result {

		if len(args) < 1 {
			return newResult(0)
		}

		args = GetArgs(args)
		var ri int64
		var rf float64
		isInt := true

		if vi, ok := args[0].Value.(int64); ok {
			ri = vi
		} else if vf, ok := args[0].Value.(float64); ok {
			isInt = false
			rf = vf
		}

		for _, v := range args[1:] {
			if vi, ok := v.Value.(int64); ok {
				if isInt {
					ri *= vi
				} else {
					rf *= float64(vi)
				}
				continue
			} else {
				rf = float64(ri)
				isInt = false
			}
			if vf, ok := v.Value.(float64); ok {
				rf *= vf
			}
		}

		if isInt {
			return newResult(ri)
		}
		return newResult(rf)
	}))

	Register("/", Function(func(args []*Result) *Result {

		if len(args) < 1 {
			return newResult(0)
		}

		args = GetArgs(args)
		var rf float64

		if vi, ok := args[0].Value.(int64); ok {
			rf = float64(vi)
		} else if vf, ok := args[0].Value.(float64); ok {
			rf = vf
		}

		for _, v := range args[1:] {
			if vi, ok := v.Value.(int64); ok {
				if vi == 0 {
					return newErrorResult(fmt.Errorf("division by zero"))
				}
				rf /= float64(vi)
				continue
			} else if vf, ok := v.Value.(float64); ok {
				if vf == 0 {
					return newErrorResult(fmt.Errorf("division by zero"))
				}
				rf /= vf
			}
		}

		return newResult(rf)
	}))

}
