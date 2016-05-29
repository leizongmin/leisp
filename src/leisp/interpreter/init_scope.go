// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"leisp/parser"
	"leisp/types"
)

var initLeispPrograms = `

;; type assertion
(defn null? [a] (equal? (typeof a) "null"))
(defn integer? [a] (equal? (typeof a) "integer"))
(defn float? [a] (equal? (typeof a) "float"))
(defn number? [a] (or (equal? (typeof a) "float") (equal? (typeof a) "integer")))
(defn boolean? [a] (equal? (typeof a) "boolean"))
(defn any? [a] (equal? (typeof a) "any"))
(defn array? [a] (equal? (typeof a) "array"))
(defn expression? [a] (equal? (typeof a) "expression"))
(defn function? [a] (equal? (typeof a) "function"))
(defn lambda? [a] (equal? (typeof a) "lambda"))
(defn infinity? [a] (equal? (typeof a) "infinity"))
(defn keyword? [a] (equal? (typeof a) "keyword"))
(defn symbol? [a] (equal? (typeof a) "symbol"))
(defn ratio? [a] (equal? (typeof a) "ratio"))
(defn string? [a] (equal? (typeof a) "string"))
(defn scope? [a] (equal? (typeof a) "scope"))

;; stdout
(defn print [a] (stdout-print (to-string a)))
(defn println [a] (print a) (stdout-print "
"))

`

func NewDefaultScope() (*types.Scope, error) {

	s := types.NewRootScope()

	RegisterBuiltinFunction(s, "lambda", builtinLambda)
	RegisterBuiltinFunction(s, "defn", builtinDefn)
	RegisterBuiltinFunction(s, "eval", builtinEval)
	RegisterBuiltinFunction(s, "func-call", builtinFunctionCall)
	RegisterBuiltinFunction(s, "func-apply", builtinFunctionApply)

	RegisterBuiltinFunction(s, "typeof", builtinTypeOf)
	RegisterBuiltinFunction(s, "defvar", builtinDefvar)
	// RegisterBuiltinFunction(s, "value", builtinValue)
	RegisterBuiltinFunction(s, "getvar", builtinValueByName)
	RegisterBuiltinFunction(s, "to-string", builtinToString)

	RegisterBuiltinFunction(s, "new-scope", builtinNewScope)
	RegisterBuiltinFunction(s, "exit", builtinExit)

	RegisterBuiltinFunction(s, "equal?", builtinEqual)
	RegisterBuiltinFunction(s, "=", builtinEqual)

	RegisterBuiltinFunction(s, "and", builtinAnd)
	RegisterBuiltinFunction(s, "or", builtinOr)
	RegisterBuiltinFunction(s, "not", builtinNot)

	RegisterBuiltinFunction(s, "array-length", builtinArrayLength)
	RegisterBuiltinFunction(s, "array-index", builtinArrayIndex)
	RegisterBuiltinFunction(s, "array-slice", builtinArraySlice)

	RegisterBuiltinFunction(s, "stdout-print", builtinStdOutPrint)

	RegisterBuiltinFunction(s, "+", builtinMathAdd)
	RegisterBuiltinFunction(s, "-", builtinMathSubtract)
	RegisterBuiltinFunction(s, "*", builtinMathMultiply)
	RegisterBuiltinFunction(s, "/", builtinMathDivide)
	RegisterBuiltinFunction(s, "^", builtinMathPow)

	RegisterBuiltinFunction(s, "str", builtinStr)

	astList, err := parser.ParseAll(initLeispPrograms)
	if err != nil {
		return nil, err
	}

	r := EvalASTList(s, astList)
	if r.IsError() {
		return nil, r.Error
	}

	return s, nil
}
