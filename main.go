package main

import (
	"portable/scanner"
)

func main() {
	scanner.NewScan(`// 变量定义
	// 不指定变量类型默认为any（任意类型
	var name = "test"
	// 指定变量类型，限定声明
	var name:number = 123
	// 常量，初始赋值后不可变
	const name:function = fn()int{
		ret 0
	}
	// 函数定义
	fn v(){
		...
	}
	//引入多个源文件
	link "demo" "fmt"
	// 引入源文件中的某个(函数/类/变量)
	link "xxxx" > "xx"
	// 引入源文件中的多个(函数/类/变量)
	link "xxxx" > "xx","xxe"
	// 注释
	// 单行注释
	//
	// 多行注释
	/**/
	
	// 类，公开访问权限
	class name{
		// 加enum就是枚举，默认从0开始枚举，可加=
		enum EnumType{
			// 枚举的变量拥有其对应的字符串形式的变量名称，可手动得到，枚举的类型则单独作一种类型，需要以var name:EnumType声明存储（可不加类型
			// 枚举变量可强转至其定义的类型，如果类型错误，则返回nil
			A,B,C,D=1
		}
		// 类中的函数
		fn MyFun(x:number,y:string,z:any)string,number{
			// 函数执行结束后执行的函数，可以在任何作用域中用
			// 如果在类作用域或类作用域内的自定义作用域中使用
			defer fn(){
				
			}
			// 返回值，支持多返回值
			ret x,y
		}
		// 构造函数
		fn name(){
	
		}
	}
	// 类泛型
	class<T> cc{
		fn cc(v:T){
			print(v)
		}
	}
	// 函数泛型
	fn<T,V> x(v:T,s:V){
		print(v,s)
	}
	// 错误捕获
	try{
		...
	}except e{
		...
	}
	//错误向上抛出
	fn x(){
		try{
			...
		}catch e{
			throw(e)
		}
	}
	
	// 主函数入口
	fn main(){
		
	}
	
	// 关于继承
	// 类class的继承都将其在其内部的所有标识符作为继承的对象即
	class k{
		
	}
	class name{
		k
	}
	//      ...
	// name2为继承的类，如果第一个表达式不是已有的类或者是并不是一个标识符而是一个其他表达式，将报错或者不会继承，如
	// (name2)
	// var k = name2
	// ......
	// 类全局变量只能通过析构函数执行
	class name{
		// 仅可继承一个类，也就是单继承
		name2
		// 构造函数不能有返回值
		fn name(b:string){
			// 调用父类中的构造函数
			self.super()
			// self代表类中的全局变量
			var self.s:int = ""
		}
			
	}
	// 一般表达为暂时未写，也就是TODO的含义，没有任何作用
	...
	// 运算符/操作符
	! & and | or ^ % ?: ?? ?. ??=
	> < >= <= != ==
	// 左值是否包含右值的内容
	in
	++ --
	^= %= += -= /= *=
	= + - * / ( ) { }
	// 切片，仅用于list
	// arr[1:2]
	
	// 说明，字节类型和字符类型不是一种类，字节代表的是1字节
	// 字符代表的是UTF8中一个字符（在英文下作用相同）虽然如此，但是byte和char同用char类型表示（byte可以用数字表示，char不行
	// 字节类型
	byte
	// 字符类型
	char
	// 字符串是字符类型的集合
	string
	// 字符串支持的转义字符（后续有需求再增加）
	// \\ 反斜杠符号
	// \' 单引号
	// \" 双引号
	// \n 换行
	// \t 横向制表符(Tab)
	
	
	// 数字类型，直接用最大数字类型表示
	number
	// 函数类型
	function
	// 可容纳任意类型
	any
	// 如果不为其定义类型则默认为any，可以存放任何值，否则仅可以存放某一类型的值
	// 如var k:table<string> = ["a","b","c"]
	list<type>
	
	// 如果不为其定义类型则默认为any，可以存放任何值，否则仅可以存放某一类型的值
	// 如var k:dict<string,string> = {a:"a",b:"b",c:"c"}
	dict<type>
	// 布尔型变量，只有true和false
	bool
	
	// 布尔值
	true
	false
	// 空，没有的意思，变量声明如果不赋予初始值，则为nil，
	nil
	// 指向类自身的指针
	self
	// 类，同结构体
	class
	// 变量声明
	var
	// 链接其他源文件
	link
	// 函数声明
	fn
	// 返回
	ret
	// 常量
	const
	// 函数结束时执行
	defer
	
	
	
	// 判断
	if
	// 否则
	else
	// 多条件判断
	switch
	// 选项
	case
	// 循环
	for
	// 退出循环或者switch
	break
	// 跳过本次循环
	continue
	// 删除字典中的键
	// del varName
	del
	// 用于为任何东西定义别名，会做预处理操作，即使是 define hello "hello" 如此也是可以的，但请小心使用
	// define hello(x) fn(){fnc(x)}
	// 
	// print(hello) -> print("hello")
	define
	// 用于取消define的定义
	// define hello "hello"
	// undef hello
	undef
	// 枚举
	enum
	
	
	
	
	//条件语句
	if(exp){
		...
	}elif(exp){
		...
	}else{
		...
	}
	// 默认不需要break，会逐条检测，一旦匹配将停止检查分支
	// switch的作用域之间是共享的，如果不希望共享请主动添加作用域{}
	switch(val){
		case a:{
			...
		}
		case b,c:
		case:
	}
	// while
	for(exp){
		
	}
	// foreach
	for(itme:list){
		
	}
	// for
	for(exp;exp;exp){
		
	}
	
	// 语法糖
	// 三元运算符
	// 同等于
	// if(exp) v1 else v2
	exp?v1:v2
	
	`).Start().Output()
}
