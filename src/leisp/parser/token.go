// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package parser

type token int

const (
	tokenIllegal token = iota
	tokenEOF
	tokenWhitespace
	tokenComment

	tokenKeyword
	tokenSymbol
	tokenNumber
	tokenString

	tokenPunctuation
	tokenQuote
)
