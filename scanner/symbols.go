package scanner

const (
	StrType     TokenType = iota // 字符串类型 "str"
	BoolType                     // 布尔类型 true false
	NullType                     // 空类型 nil
	ByteType                     // 字节类型 uint8，表示类似于数字，但最大为255
	CharType                     // 字符类型 'X'
	NumberType                   // 数字类型 123 1.23
	CommentType                  // 注释类型 //xx /*xxx*/
	ListType                     // 列表类型 [1,23,44]
	DictType                     // 字典 {1:2,"a":b}

	DotOperator            // 点运算符 .
	ExclamOperator         // 逻辑非操作符 !
	AmpersandOperator      // 位与操作符 &
	AndOperator            // 逻辑与操作符 and
	PipeOperator           // 位或操作符 |
	OrOperator             // 逻辑或操作符 or
	CaretOperator          // 位异或操作符 ^
	PercentOperator        // 百分号操作符 %
	ContainsOperator       // in 运算符，用于判断集合中是否包含元素
	FallbackOperator       // ?? 后备运算符，返回传递的两个参数中的非空值
	AccessIfNotNilOperator // ?. 条件访问运算符，仅在操作数不是nil时才执行访问操作
	QuestionOperator       // 问号 ?
	ColonOperator          // 冒号 :
	NullCoalesceOperator   // 空合并操作符 ??
	PlusOperator           // 加号 +
	MinusOperator          // 减号 -
	MultiplyOperator       // 乘号 *
	DivideOperator         // 除号 /

	GT  // 大于号 >
	LT  // 小于号 <
	GTE // 大于等于号 >=
	LTE // 小于等于号 <=
	NEQ // 不等于号 !=
	EQ  // 等于号 ==

	Semicolon // 分号 ;
	Comma     // 逗号 ,

	LParen   // 左括号 (
	RParen   // 右括号 )
	LBrace   // 左大括号 {
	RBrace   // 右大括号 }
	LBracket // 左方括号 [
	RBracket // 右方括号 ]

	IncAssign // 自增操作符 ++
	DecAssign // 自减操作符 --

	Assign         // 赋值操作符 =
	PlusAssign     // 加等操作符 +=
	MinusAssign    // 减等操作符 -=
	MultiplyAssign // 乘等操作符 *=
	DivideAssign   // 除等操作符 /=
	ModuloAssign   // 取模等操作符 %=
	BitAndAssign   // 位与等操作符 &=
	BitOrAssign    // 位或等操作符 |=
	BitXorAssign   // 位异或等操作符 ^=
	FallbackAssign // ??= 后备赋值操作符，如果变量的值为nil，则赋予新的值

)
