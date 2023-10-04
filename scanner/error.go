package scanner

import (
	"fmt"
	"os"
	"strconv"
)

// 语法错误：意外的Token
func UnexpectedTokenError(ch ScanChar) {
	fmt.Printf("SyntaxError(%s:%s): Invalid or unexpected token", strconv.Itoa(ch.Line), strconv.Itoa(ch.Pos))
	os.Exit(0)
}

// 语法错误：应为'x'，但得到意外Token 'y'
func CaptureToUnexpectedTokenError(ch ScanChar, x rune, y rune) {
	fmt.Printf("SyntaxError(%s:%s): Expected '%s' but got unexpected token '%s'", strconv.Itoa(ch.Line), strconv.Itoa(ch.Pos), string(x), string(y))
	os.Exit(0)
}

// 语法错误：应为'x'，但得到意外Token 'y'
func CaptureToEOFError(ch ScanChar, x rune) {
	fmt.Printf("SyntaxError(%s:%s): Expected '%s' but <EOF> was encountered", strconv.Itoa(ch.Line), strconv.Itoa(ch.Pos), string(x))
	os.Exit(0)
}
