package main

import (
	"nenki/scanner"
)

func main() {
	scanner.NewScan(`// 变量定义
	1999999999999999999+ 20*10`).Start().Output()
}
