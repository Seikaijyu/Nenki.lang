package scanner

import "fmt"

type ScanChar struct {
	Ch   Char
	Pos  int
	Line int
}

// 创建一个扫描器字符
func NewScanChar(ch Char, pos int, line int) *ScanChar {
	return &ScanChar{
		Ch:   Char(ch),
		Pos:  pos,
		Line: line,
	}
}

type Char rune

func (p *Char) Output() {
	fmt.Println(string(*p))
}

// 检查字符是否为0-9之间的数字
func (p *Char) IsNumberChar() bool {
	return *p > 47 && *p < 58
}

// 检查字符是否为英文字符
func (p *Char) IsEnglishChar() bool {
	// 65-90这个区间的ASCII码是A-Z的，97-122这个区间的ASCII码是a-z的
	return (*p > 64 && *p < 91) || (*p > 96 && *p < 123)
}

// 检查字符是否为空白符
func (p *Char) IsWhiteSpaceChar() bool {
	return *p == ' ' || *p == '\n' || *p == '\t'
}
