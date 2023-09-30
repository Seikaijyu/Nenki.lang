package main

import (
	"portable/scanner"
)

func main() {
	/*var a:number= 29
	var b:number= 30
	fn add(a:number,b:number)->number{
		 ret a+b
	}
	print(add(a,b))`, "filename*/
	scanner.NewScan(`-
	
	
	+ *//*++>=<=*=/= ????=?.`, "").Start().Output()
}
