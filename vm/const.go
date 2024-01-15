package vm

type StackTYPE int

// 栈类型
const (
	NilType      StackTYPE = iota // 空类型
	ErrorType                     // 错误类型
	IntegerType                   // 整数类型
	StringType                    // 字符串类型
	BoolType                      // 布尔类型
	ListType                      // 列表类型
	DictType                      // 字典类型
	FunctionType                  // 函数类型
)
