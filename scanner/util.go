package scanner

type Char rune

// 通过ASCII码检查是不是0-9之间的数字
func (p *Char) isNumber() bool {
	return *p > 47 && *p < 58
}

// 判断是否为英文字符
func (p *Char) isEnglishChar() bool {
	// 65-90这个区间的ASCII码是A-Z的，97-122这个区间的ASCII码是a-z的
	return (*p > 64 && *p < 91) || (*p > 96 && *p < 123)
}
