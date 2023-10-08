package ast

import "nenki/scanner"

// 函数定义，包含具名函数和匿名函数
type FnExp struct {
	// 函数名字，匿名函数则为""
	FnName string
	// 形式参数
	FormalParams Params
	// 返回值类型
	RetTypes Types
	Data     ASTData
}

// 运算表达式
type OperationalExp struct {
	Operator scanner.TokenType
	Lvalue   any
	Rvalue   any
	Data     ASTData
}
