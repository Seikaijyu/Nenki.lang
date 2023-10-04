package scanner

import (
	"bytes"
)

type Scanner struct {
	code   *ScanString
	tokens *TokenArray
}

// 词法分析器，一次性完成代码的所有词法分析步骤（方便回退Token，但是词法分析器的性能会降低）
func NewScan(code string) *Scanner {
	var scan = &Scanner{
		code:   NewScanString(code),
		tokens: NewTokenArray(),
	}
	return scan
}

// 添加一个不需要值的Token
func (p *Scanner) addSymbolToken(id TokenType) {
	p.addToken(id, nil)
}

// 添加一个Token
func (p *Scanner) addToken(id TokenType, value any) {
	p.tokens.Add(&Token{
		ID:    id,
		Value: value,
		Data: &TokenData{
			Line: p.code.Line,
			Pos:  p.code.Pos,
		},
	})
}

// 寻找符号
func (p *Scanner) matchSymbol() bool {

	switch p.code.CurrChar {
	case '.':
		p.addSymbolToken(DotOperator)
	case '!':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(NEQ)
		} else {
			p.addSymbolToken(ExclamOperator)
		}
	case '&':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(BitAndAssign)
		} else {
			p.addSymbolToken(AmpersandOperator)
		}
	case '|':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(BitOrAssign)
		} else {
			p.addSymbolToken(PipeOperator)
		}
	case '^':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(BitXorAssign)
		} else {
			p.addSymbolToken(CaretOperator)
		}
	case '%':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(ModuloAssign)
		} else {
			p.addSymbolToken(PercentOperator)
		}
	case '?':
		if p.code.MatchNextAndSkip('?') {
			if p.code.MatchNextAndSkip('=') {
				p.addSymbolToken(NullConditionalAssignment)
			} else {
				p.addSymbolToken(NullCoalesceOperator)
			}
		} else if p.code.MatchNextAndSkip('.') {
			p.addSymbolToken(AccessIfNotNullOperator)
		} else {
			p.addSymbolToken(QuestionOperator)
		}
	case ':':

		p.addSymbolToken(ColonOperator)
	case '+':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(PlusAssign)
		} else if p.code.MatchNextAndSkip('+') {
			p.addSymbolToken(IncAssign)
		} else {
			p.addSymbolToken(PlusOperator)
		}
	case '-':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(MinusAssign)
		} else if p.code.MatchNextAndSkip('-') {
			p.addSymbolToken(DecAssign)
		} else {
			p.addSymbolToken(MinusOperator)
		}
	case '*':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(MultiplyAssign)
		} else {
			p.addSymbolToken(MultiplyOperator)
		}
	case '/':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(DivideAssign)
		} else {
			p.addSymbolToken(DivideOperator)
		}
	case '>':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(GTE)
		} else {
			p.addSymbolToken(GT)
		}
	case '<':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(LTE)
		} else {
			p.addSymbolToken(LT)
		}
	case ';':
		p.addSymbolToken(Semicolon)
	case ',':
		p.addSymbolToken(Comma)
	case '(':
		p.addSymbolToken(LParen)
	case ')':
		p.addSymbolToken(RParen)
	case '{':
		p.addSymbolToken(LBrace)
	case '}':
		p.addSymbolToken(RBrace)
	case '[':
		p.addSymbolToken(LBracket)
	case ']':
		p.addSymbolToken(RBracket)
	case '=':
		if p.code.MatchNextAndSkip('=') {
			p.addSymbolToken(EQ)
		} else {
			p.addSymbolToken(Assign)
		}
	default:
		return false
	}

	return true
}

// UnexpectedTokenError 的包装
func (p *Scanner) UnexpectedTokenError() {
	UnexpectedTokenError(*p.code.cache[len(p.code.cache)-1])
}

// CaptureCharFailedError 的包装
func (p *Scanner) CaptureToUnexpectedTokenError(except rune) {
	CaptureToUnexpectedTokenError(*p.code.cache[len(p.code.cache)-1], except, rune(p.code.NextChar))
}

// CaptureToEOFError 的包装
func (p *Scanner) CaptureToEOFError(except rune) {
	CaptureToEOFError(*p.code.cache[len(p.code.cache)-1], except)
}

// 捕获字符串（支持多行字符串）
func (p *Scanner) matchString() bool {

	if p.code.CurrChar == '"' {
		var temp bytes.Buffer
		if p.code.IsEndOfInput() {
			p.CaptureToEOFError('"')
		}
		for !p.code.MatchNextAndSkip('"') {
			if p.code.MatchNextAndSkip('\\') {
				temp.WriteRune(rune(p.escapeCharTrans()))
			} else {
				temp.WriteRune(rune(p.code.NextChar))
			}
			if p.code.Next() {
				p.CaptureToEOFError('"')
			}
		}
		p.addToken(StrType, temp.String())
		return true
	}

	return false
}

// 字符串内的转义字符处理
func (p *Scanner) escapeCharTrans() Char {
	switch p.code.NextChar {
	case 'n':
		return '\n'
	case 't':
		return '\t'
	default:
		return p.code.NextChar
	}
}

// 捕获字符类型
func (p *Scanner) matchChar() bool {
	if p.code.CurrChar == '\'' {
		if p.code.MatchNextAndSkip('\\') {
			p.addToken(CharType, p.escapeCharTrans())
		} else if p.code.NextChar == '\'' {
			p.UnexpectedTokenError()
		} else {
			p.addToken(CharType, rune(p.code.NextChar))
		}
		if p.code.Next() {
			p.CaptureToEOFError('\'')
		}
		if p.code.MatchNextAndSkip('\'') {
			return true
		}
		p.CaptureToUnexpectedTokenError('\'')

	}
	return false
}

// 捕获数字（包括浮点数）
func (p *Scanner) matchNumber() bool {

	if p.code.CurrChar.IsNumberChar() {
		var temp bytes.Buffer
		// 是否为浮点数
		var isFloatNumber bool = false
		temp.WriteRune(rune(p.code.CurrChar))
		for {
			if p.code.IsEndOfInput() {
				break
			}
			if p.code.NextChar.IsNumberChar() {
				temp.WriteRune(rune(p.code.NextChar))
			} else if p.code.NextChar == '.' {
				if isFloatNumber {
					p.UnexpectedTokenError()
				} else {
					temp.WriteRune(rune(p.code.NextChar))
					isFloatNumber = true
				}
			} else {
				break
			}
			p.code.Next()
		}
		if isFloatNumber {
			if temp.Bytes()[temp.Len()-1] != '.' {
				p.addToken(NumberType, NewNumber(isFloatNumber, temp.String()))
			} else {
				p.UnexpectedTokenError()
			}
		} else {
			p.addToken(NumberType, NewNumber(isFloatNumber, temp.String()))
		}
		return true
	}
	return false
}

// 跳过空白字符
func (p *Scanner) skipWhiteSpace() {
	for p.code.CurrChar.IsWhiteSpaceChar() {
		if p.code.Next() {
			return
		}
	}
}

// 捕获注释
func (p *Scanner) matchComment() bool {
	if p.code.CurrChar == '/' {
		var temp bytes.Buffer

		if p.code.MatchNextAndSkip('/') {
			for p.code.NextChar != '\n' {
				temp.WriteRune(rune(p.code.NextChar))
				if p.code.Next() {
					break
				}
			}
			p.addToken(CommentType, temp.String())
			return true
		} else if p.code.MatchNextAndSkip('*') {
			for {
				if p.code.MatchNextAndSkip('*') {
					if p.code.MatchNextAndSkip('/') {
						break
					}
					temp.WriteRune('*')
				} else {
					temp.WriteRune(rune(p.code.NextChar))
				}
				if p.code.Next() {
					p.UnexpectedTokenError()
				}
			}
			p.addToken(CommentType, temp.String())
			return true
		}
	}
	return false
}

// 查找标识符符号（在标识符之中有一部分是符号）
func (p *Scanner) matchIdentSymbol(ident string) bool {
	switch ident {
	case "in":
		p.addSymbolToken(ContainsOperator)
	case "and":
		p.addSymbolToken(AndOperator)
	case "or":
		p.addSymbolToken(OrOperator)
	default:
		return false
	}
	return true
}

// 捕获标识符内的各类型值
func (p *Scanner) matchIdentTypeValue(ident string) bool {
	switch ident {
	case "true":
		p.addToken(BoolType, true)
	case "false":
		p.addToken(BoolType, false)
	case "nil":
		p.addSymbolToken(NullType)
	default:
		return false
	}
	return true
}

// 捕获保留关键字
func (p *Scanner) matchIdentKeyWord(ident string) bool {
	for _, v := range keywords {
		if ident == v {
			p.addToken(KeyWord, v)
			return true
		}
	}
	return false
}

// 捕获标识符，支持a-zA-Z_$0-9
func (p *Scanner) matchIdentifier() bool {
	if p.code.CurrChar.IsIdentifierChar() {
		var temp bytes.Buffer
		temp.WriteRune(rune(p.code.CurrChar))
		for {
			if p.code.IsEndOfInput() {
				break
			}
			if p.code.NextChar.IsIdentifierChar() {
				temp.WriteRune(rune(p.code.NextChar))
			} else if p.code.NextChar.IsWhiteSpaceChar() ||
				p.code.NextChar.FromListFind([]rune{'+', '-', '*', '/', '^', '?', '.', '!', '=', '&', '|', '%', ':', '>', '<', ';', ',', '(', ')', '[', ']', '{', '}'}) {
				break
			} else {
				p.UnexpectedTokenError()
			}
			p.code.Next()
		}
		identifier := temp.String()
		if p.matchIdentSymbol(identifier) {
		} else if p.matchIdentTypeValue(identifier) {
		} else if p.matchIdentKeyWord(identifier) {
		} else {
			p.addToken(Identifier, temp.String())
		}

		return true
	}
	return false
}

func (p *Scanner) Start() *TokenArray {
	for {
		p.skipWhiteSpace()
		if p.code.IsEndOfInput() {
			return p.tokens
		}
		if ok := p.matchComment(); ok {
		} else if ok := p.matchSymbol(); ok {
		} else if ok := p.matchString(); ok {
		} else if ok := p.matchChar(); ok {
		} else if ok := p.matchNumber(); ok {
		} else if ok := p.matchIdentifier(); ok {
		} else {
			UnexpectedTokenError(*p.code.cache[len(p.code.cache)-1])
		}
		p.code.Next()

	}

}
