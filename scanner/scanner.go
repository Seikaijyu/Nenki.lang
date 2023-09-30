package scanner

type Scanner struct {
	code     *ScanString
	tokens   *TokenArray
	fileName string
}

// 词法分析器，一次性完成代码的所有词法分析步骤（方便回退Token，但是词法分析器的性能会降低）
func NewScan(code string, fileName string) *Scanner {
	var scan = &Scanner{
		code:     NewScanString(code),
		tokens:   NewTokenArray(),
		fileName: fileName,
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
			Name: p.fileName,
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

// 跳过空白字符
func (p *Scanner) skipWhiteSpace() {
	for p.code.CurrChar.IsWhiteSpaceChar() {
		if p.code.Next() {
			return
		}
	}
}

func (p *Scanner) Start() *TokenArray {
	for {
		p.skipWhiteSpace()
		if ok := p.matchSymbol(); ok {

		}
		if p.code.Next() {
			return p.tokens
		}
	}

}
