package scanner

type ScanString struct {
	str   []rune
	index int
	Pos   int
	Line  int
	// 上文缓存
	cache []*ScanChar
	// 当前字符
	CurrChar Char
	// 下一个字符
	NextChar Char
}

// 创建一个扫描器字符串
func NewScanString(v string) *ScanString {

	var scanString = &ScanString{
		str:   []rune(v),
		index: 0,
		Pos:   0,
		Line:  1,
	}
	scanString.Next()

	return scanString
}

// 寻找下一个字符是否与匹配字符相等，相等则返回true并Next
func (p *ScanString) MatchNextAndSkip(ch Char) bool {
	if p.NextChar == ch {
		p.Next()
		return true
	}
	return false
}

// 获取下一个字符
func (p *ScanString) Next() bool {
	if p.IsEndOfInput() {
		p.NextChar = 0
		return true
	}
	// 如果当前索引数量大于缓存数量
	if p.index > len(p.cache)-1 {
		// 缓存上一个扫描好的信息
		p.cache = append(p.cache, NewScanChar(p.CurrChar, p.Pos, p.Line))
		// 只保留数组中最新的4条数据以保证不会占用过多内存
		if len(p.cache) > 4 {
			// 移动元素，保持数组中最新的 'maxSize' 个元素
			p.cache = p.cache[len(p.cache)-4:]
		}
	} else {
		p.CurrChar = p.cache[p.index].Ch
		p.Pos = p.cache[p.index].Pos
		p.Line = p.cache[p.index].Line
		p.index++
		return p.IsEndOfInput()
	}

	p.CurrChar = Char(p.str[p.index])

	if p.CurrChar == '\n' {
		p.Pos = 0
		p.Line++
	} else {
		p.Pos++
	}
	p.index++
	if !p.IsEndOfInput() {
		p.NextChar = Char(p.str[p.index])
	}

	return false
}

// 回溯到上一个
func (p *ScanString) RetToPrevious() {
	if p.index <= 1 {
		return
	}
	p.index--

	p.NextChar = Char(p.str[p.index])

	p.CurrChar = p.cache[p.index].Ch
	p.Pos = p.cache[p.index].Pos
	p.Line = p.cache[p.index].Line
}

// 检查是否到底了
func (p *ScanString) IsEndOfInput() bool {
	return p.index >= len(p.str)
}
