package scanner

type Scanner struct {
	tokens   *TokenArray
	fileName string
	pos      DataInt
	line     DataInt
	tempChar Char
}

// 词法分析器，一次性完成代码的所有词法分析步骤（方便回退Token，但是词法分析器的性能会降低）
func NewScan(fileName string) *Scanner {
	var scan = &Scanner{
		tokens:   NewTokenArray(),
		fileName: fileName,
	}
	return scan
}

func (p *Scanner) Start() *TokenArray {
	return p.tokens
}
