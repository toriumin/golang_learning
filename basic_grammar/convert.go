package main

import "fmt"

func main() {

	var i int = 1234

	var u uint32 = uint32(i)

	var f float32 = float32(u)

	var s string = string(i)

	var b []byte = []byte("abc")

	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
}

// 出力結果
// 1234
// 1234
// Ӓ
// [97 98 99]
