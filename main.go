package main

import (
	"portable/scanner"
)

func main() {
	scanner.NewScan(`// 变量定义
	name.a
	`).Start().Output()
}
