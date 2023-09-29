### <p style="text-align:center;">Portable语言关键字运算符预览</p>
<br/>
###### 关键字（Keywords）
| 关键字 | 说明 |
|---|---|
|var|用于声明变量|
|const|用于声明常量|
|ret|函数返回值关键字，与return作用相同|
|link|导入其他包或者包中的函数，变量等|
|class|用于类的声明|
|enum|声明枚举常量|
|defer|函数执行结束后执行，相当于在函数返回值之前执行的函数|
|try|捕获错误|
|catch|处理捕获的错误|
|if|逻辑判断|
|elif|否则继续进行逻辑判断|
|else|判断都不成功执行|
|swicth|多条件分支|
|case|分支选项，如果没有值则作为其他分支未命中执行的分支|
|for|循环，可以迭代值|
|break|退出循环|
|continue|单次跳过循环|
|del|删除字典中的键|
|define|预处理替换，可以为类型定义别名|
|undef|取消预处理替换|
|byte|字节类型，仅表示单个字节|
|char|字符类型，仅表示utf-8编码字符|
|string|字符串类型，表示一组char字符|
|number|数字类型，可以为浮点数和整数|
|function|函数类型，可以存储任何函数|
|any|任何类型，可以容纳任何类型的值，但是使用前必须转换类型|
|list|列表类型，一组连续空间的列表|
|dict|字典类型，表示一组k-v字典|
|bool|布尔类型，只有true和false|
|nil|空类型，没有任何值|
|true|布尔真值|
|false|布尔假值|
|self|表示自身，仅在函数和类中可用，如果函数在类中则表示类本身|


<br/>
###### 操作符（Symbols）
| 操作符 | 一般操作数  | 使用例子 | 其他说明 |
|---|---|---|---|
|!|1|!a|
|++|1|a++ ++a|
|--|1|a-- --a|
|.|2|a.b|
|&|2|a&b|
|\||2|a\|b|
|^|2|a^b|
|%|2|a%b|
|+|2|a+b|
|-|2|a-b|
|*|2|a*b|
|/|2|a/b|
|=|2|a=b|
|&=|2|a&=b|
|\|=|2|a\|=b|
|^=|2|a^=b|
|%=|2|a%=b|
|+=|2|a+=b|
|-=|2|a-=b|
|*=|2|a*=b|
|/=|2|a/=b|
|??=|2|a??=b|如果a的值为nil则a=b
|in|2|a in b|判断左值是否包含右值的内容
|>|2|a>b|
|<|2|a<b|
|>=|2|a>=b|
|<=|2|a<=b|
|\=\=|2|a\=\=b|
|\!\=|2|a\!\=b|
|??|2|a??b|a的值如果为nil则返回b
|?.|2|a?.b()|a中如果没有b则不会执行
|and|2|a\=\=1 and b\=\=2|替代&&
|or|2|a\=\=1 or b\=\=2|替代\|\|
|?:|3|a?b:c|
