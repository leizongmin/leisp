// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type Token int

const (
	TokenIllegal Token = iota
	TokenEOF
	TokenWhitespace
	TokenComment

	TokenSymbol
	TokenNumber
	TokenString
	TokenChar

	TokenPunctuation
)
