package basic_grammar

import "fmt"

func main() {

	var a int = 12345

	// int 型の変数 a を int64 型に変換する
	var i64 int64 = int64(a)

	fmt.Println(i64)
}
