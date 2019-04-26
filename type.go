package main

import "fmt"

func main() {

	// int 型から myInteger 型を宣言する
	type myInteger int

	var a myInteger = 123

	a += 1

	fmt.Println(a)

	//
	type myStruct struct {
		a int
		b int
	}
}

// 出力結果
// 124
