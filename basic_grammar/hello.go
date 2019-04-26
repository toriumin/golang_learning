package basic_grammar

// main パッケージの中にある fmt 関数
// 使う関数しかインポートしちゃだめ
import "fmt"

// fmt パッケージ内に main 関数がある
// main は必ず一番最初に読み込まれる関数
func main() {
	// p は大文字
	// なんのモジュールを使ってるか書く
	// 外部からパッケージをインポートした場合、パッケージと関数の間にピリオドを入れなければならない
	fmt.Println("Hello Golang!")
}
