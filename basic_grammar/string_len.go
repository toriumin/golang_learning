package basic_grammar

import "fmt"

func main() {

	var en string = "golang"
	var ja string = "Go言語"

	fmt.Println(en, " len:", len(en))
	fmt.Println(ja, " len:", len(ja))
}

//出力結果
// golang len:6
// Go言語 len:8
