package scanner

var keywords []string = []string{
	// 关键字
	"var",      // 变量
	"const",    // 常量
	"fn",       // 函数声明
	"ret",      // 返回值
	"link",     // 导入包
	"class",    // 类
	"enum",     // 枚举
	"defer",    // 函数结束后执行
	"try",      // 错误捕获
	"catch",    // 错误处理
	"if",       // 判断
	"elif",     // 否则判断
	"else",     // 否则
	"swicth",   // 多条件分支
	"case",     // 分支选项
	"for",      // 循环
	"break",    // 退出循环
	"continue", // 跳过单次循环
	"del",      // 删除字典中的键
	"define",   // 预处理定义
	"undef",    // 取消预处理定义

	// 值
	"true",  // 布尔是
	"false", // 布尔否
	"nil",   // 空值（或无值）
	"self",  // 类自身

	// 类型
	"byte",     // 字节类型
	"char",     // 字符类型
	"string",   // 字符串类型
	"number",   // 数字类型
	"function", // 函数类型
	"any",      // 任何类型
	"list",     // 列表类型
	"dict",     // 字典类型
	"bool",     // 布尔类型
}

// 用于查找键是否存在，返回一个布尔值
// 成功返回true
func FindKeyWord(find string) bool {
	for _, value := range keywords {
		if find == value {
			return true
		}
	}
	return false
}
