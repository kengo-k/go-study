// 最初にパッケージ名がくる
// mainパッケージは特別なパッケージで実行形式毎に必ず含まれていなければならない
package main

import (
	"fmt"
)

// mainパッケージのmain関数がエントリポイントになる
func main() {
	fmt.Println("Hello World")
}
