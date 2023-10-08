package ast

type Param struct {
	// 变量类型
	Type string
	// 变量名称
	Name string
	// 起始行
	Data ASTData
}

type Type struct {
	// 变量类型
	Type string
	Data ASTData
}
type ASTData struct {
	Line int
	Pos  int
}
type Types []Type
type Params []Param
