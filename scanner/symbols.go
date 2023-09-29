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

	Scope        // 作用域（在语言中代表代码之前空格的数量，最大255个使用uint8表示）
	Dot          // 点运算符 .
	Exclam       // 逻辑非操作符 !
	Ampersand    // 位与操作符 &
	And          // 逻辑与操作符 and
	Pipe         // 位或操作符 |
	Or           // 逻辑或操作符 or
	Caret        // 位异或操作符 ^
	Percent      // 百分号操作符 %
	Question     // 问号 ?
	Colon        // 冒号 :
	NullCoalesce // 空合并操作符 ??
	Inc          // 自增操作符 ++
	Dec          // 自减操作符 --
	Assign       // 赋值操作符 =

	GT  // 大于号 >
	LT  // 小于号 <
	GTE // 大于等于号 >=
	LTE // 小于等于号 <=
	NEQ // 不等于号 !=
	EQ  // 等于号 ==

	Semicolon // 分号 ;
	Comma     // 逗号 ,

	Plus     // 加号 +
	Minus    // 减号 -
	Multiply // 乘号 *
	Divide   // 除号 /

	LParen   // 左括号 (
	RParen   // 右括号 )
	LBrace   // 左大括号 {
	RBrace   // 右大括号 }
	LBracket // 左方括号 [
	RBracket // 右方括号 ]
	Arrow    // 箭头操作符 ->

	PlusAssign     // 加等操作符 +=
	MinusAssign    // 减等操作符 -=
	MultiplyAssign // 乘等操作符 *=
	DivideAssign   // 除等操作符 /=
	ModuloAssign   // 取模等操作符 %=
	BitAndAssign   // 位与等操作符 &=
	BitOrAssign    // 位或等操作符 |=
	BitXorAssign   // 位异或等操作符 ^=
)
