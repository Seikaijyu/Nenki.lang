package ast

// Param 表示函数或方法的参数
type Param struct {
	// 变量类型
	Type string
	// 变量名称
	Name string
	// 起始行
	Data ASTData
}

// Type 表示变量类型
type Type struct {
	// 变量类型
	Type string
	// 起始行
	Data ASTData
}

// ASTData 表示抽象语法树节点的位置信息
type ASTData struct {
	// 行号
	Line int
	// 列号
	Pos int
}

// Types 表示多个变量类型
type Types []Type

// Params 表示多个函数或方法的参数
type Params []Param
