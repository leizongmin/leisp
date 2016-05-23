// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package parser

import (
	"fmt"
	"io"
	"leisp/types"
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

func (p *Parser) Parse() (*types.AST, error) {

	tok, lit := p.scanIgnoreWhitespaceOrComment()
	switch tok {
	case TokenChar:
		return types.NewValueAST(types.NewChar(lit)), nil
	case TokenString:
		return types.NewValueAST(types.NewString(lit)), nil
	case TokenEOF:
		return types.NewEOFAST(), nil
	case TokenNumber:
		return p.parseNumber(lit)
	case TokenKeyword:
		return types.NewValueAST(types.NewKeyword(lit)), nil
	case TokenSymbol:
		LIT := strings.ToUpper(lit)
		if LIT == "NIL" {
			return types.NewValueAST(types.NewNull()), nil
		} else if LIT == "T" {
			return types.NewValueAST(types.NewBoolean(true)), nil
		} else {
			return types.NewValueAST(types.NewSymbol(lit)), nil
		}
	case TokenPunctuation:
		return p.parsePunctuation(lit)
	default:
		return types.NewEOFAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseNumber(lit string) (*types.AST, error) {

	if i := strings.IndexAny(lit, "/"); i != -1 {
		return types.NewValueAST(types.NewRatio(lit)), nil
	}
	if i := strings.IndexAny(lit, "."); i != -1 {
		if val, err := strconv.ParseFloat(lit, 64); err != nil {
			return types.NewEOFAST(), fmt.Errorf("invalid float number %s", lit)
		} else {
			return types.NewValueAST(types.NewFloat(val)), nil
		}
	}
	if val, err := strconv.ParseInt(lit, 10, 64); err != nil {
		return types.NewEOFAST(), fmt.Errorf("invalid integer number %s", lit)
	} else {
		return types.NewValueAST(types.NewInteger(val)), nil
	}
}

func (p *Parser) parsePunctuation(lit string) (*types.AST, error) {

	switch lit {
	case "(":
		return p.parseSExpression(lit)
	case "{":
		return p.parseQExpression(lit)
	case "[":
		return p.parseList(lit)
	case ")", "}", "]":
		return types.NewEOFAST(), fmt.Errorf("mismatching %s", lit)
	default:
		return types.NewEOFAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseSExpression(lit string) (*types.AST, error) {

	var children []*types.AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == TokenPunctuation && lit == ")" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return types.NewEOFAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return types.NewSExpressionAST(children), nil
}

func (p *Parser) parseQExpression(lit string) (*types.AST, error) {

	var children []*types.AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == TokenPunctuation && lit == "}" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return types.NewEOFAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return types.NewQExpressionAST(children), nil
}

func (p *Parser) parseList(lit string) (*types.AST, error) {

	var children []*types.AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == TokenPunctuation && lit == "]" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return types.NewEOFAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return types.NewListAST(children), nil
}
