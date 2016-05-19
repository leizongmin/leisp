// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
	}
}

func (p *Parser) scan() (tok Token, lit string) {

	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()
	p.buf.tok, p.buf.lit = tok, lit

	return tok, lit
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWhitespaceOrComment() (tok Token, lit string) {
	for {
		tok, lit = p.scan()
		if tok != TokenWhitespace && tok != TokenComment {
			break
		}
	}
	return tok, lit
}

func (p *Parser) GetPosition() Position {
	return p.s.Position
}

func (p *Parser) Parse() (*AST, error) {

	tok, lit := p.scanIgnoreWhitespaceOrComment()
	switch tok {
	case TokenChar:
		return newCharAST(lit), nil
	case TokenString:
		return newStringAST(lit), nil
	case TokenEOF:
		return newEmptyAST(), nil
	case TokenNumber:
		return p.parseNumber(lit)
	case TokenKeyword:
		return newKeywordAST(lit), nil
	case TokenSymbol:
		LIT := strings.ToUpper(lit)
		if LIT == "NIL" {
			return newNullAST(), nil
		} else if LIT == "T" {
			return newBooleanAST(), nil
		} else {
			return newSymbolAST(lit), nil
		}
	case TokenPunctuation:
		return p.parsePunctuation(lit)
	default:
		return newEmptyAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseNumber(lit string) (*AST, error) {

	if i := strings.IndexAny(lit, "/"); i != -1 {
		return newRatioAST(lit), nil
	}
	if i := strings.IndexAny(lit, "."); i != -1 {
		if val, err := strconv.ParseFloat(lit, 64); err != nil {
			return newEmptyAST(), fmt.Errorf("invalid float number %s", lit)
		} else {
			return newFloatAST(val), nil
		}
	}
	if val, err := strconv.ParseInt(lit, 10, 64); err != nil {
		return newEmptyAST(), fmt.Errorf("invalid integer number %s", lit)
	} else {
		return newIntegerAST(val), nil
	}
}

func (p *Parser) parsePunctuation(lit string) (*AST, error) {

	switch lit {
	case "(":
		return p.parseSExpression(lit)
	case "{":
		return p.parseQExpression(lit)
	case "[":
		return p.parseList(lit)
	case ")", "}", "]":
		return newEmptyAST(), fmt.Errorf("mismatching %s", lit)
	default:
		return newEmptyAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseSExpression(lit string) (*AST, error) {

	var children []*AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == TokenPunctuation && lit == ")" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return newEmptyAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return newSExpressionAST(children), nil
}

func (p *Parser) parseQExpression(lit string) (*AST, error) {

	var children []*AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == TokenPunctuation && lit == "}" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return newEmptyAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return newQExpressionAST(children), nil
}

func (p *Parser) parseList(lit string) (*AST, error) {

	var children []*AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == TokenPunctuation && lit == "]" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return newEmptyAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return newListAST(children), nil
}
