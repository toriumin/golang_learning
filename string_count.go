package main

import "fmt"
import "unicode/utf8"

func main() {

	var ja string = "Go言語"

	// 文字数を出力
	// rune の実体は int32 で，内部表現は Unicode の符号点 (code point)
	// string と rune 配列は相互変換できる
	fmt.Println(ja, "len:", utf8.RuneCountInString(ja))

}

// 出力結果
// Go言語 len: 4
