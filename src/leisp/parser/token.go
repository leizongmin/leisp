// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package parser

type Token int

const (
	TokenIllegal Token = iota
	TokenEOF
	TokenWhitespace
	TokenComment

	TokenKeyword
	TokenSymbol
	TokenNumber
	TokenString
	TokenChar

	TokenPunctuation
)
