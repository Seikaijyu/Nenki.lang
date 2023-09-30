package scanner

import "fmt"

type TokenType byte

// 存储Token包含的信息
type TokenData struct {
	// Token存在的链接文件名
	Name string
	// Token出现的首行
	Line int
	// Token出现的首位置
	Pos int
}

// 存储Token
type Token struct {
	// 记录Token的类型ID
	ID TokenType
	// Token的额外数据
	Data *TokenData
	// 值
	Value any
}

type TokenArray struct {
	tokens    *[]*Token
	preToken  *Token
	nextToken *Token
}

// 创建一个TokenArray类型
func NewTokenArray() *TokenArray {
	return &TokenArray{
		tokens:    &[]*Token{},
		preToken:  nil,
		nextToken: nil,
	}
}

// 把值传递到数组中中
func (p *TokenArray) Add(token *Token) {
	*(p.tokens) = append(*(p.tokens), token)
}

// 从索引提取值
func (p *TokenArray) FromIndex(id int) *Token {
	if p.Length()-1 > 0 {
		if id > 0 {
			p.preToken = (*p.tokens)[id-1]
		} else {
			p.preToken = nil
		}

		if id < p.Length()-1 {
			p.nextToken = (*p.tokens)[id+1]
		} else {
			p.nextToken = nil
		}
	}
	return (*p.tokens)[id]
}

// 返回长度
func (p *TokenArray) Length() int {
	return len(*p.tokens)
}

// 返回从索引引用的上一个Token
func (p *TokenArray) Previous() *Token {
	return p.preToken
}

// 返回从索引引用的下一个Token
func (p *TokenArray) Next() *Token {
	return p.nextToken
}

// 输出所有值
func (p *TokenArray) Output() {
	for _, token := range *p.tokens {

		fmt.Printf("ID:%s,Value:%v,Name:%s,Len:%d,Pos:%d\n",
			tokenTypeMap[token.ID], token.Value, token.Data.Name, token.Data.Line, token.Data.Pos)
	}
}
