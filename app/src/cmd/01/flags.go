package main

import (
	"flag"
	"log"
)

var (
	// 第一引数: オプション名
	// 第二引数: デフォルト値
	// 第三引数: 説明
	//
	// 引数は以下のように指定する
	// -s hello
	// -s=hello
	// -s="hello" -i 200
	FlagStr = flag.String("s", "default", "文字列フラグ")
	FlagInt = flag.Int("i", -1, "数値フラグ")
)

func main() {
	// 想定外のフラグを渡すとusageを表示して処理を終了する
	flag.Parse()
	log.Println(*FlagStr)
	log.Println(*FlagInt)
	// Argsを呼び出すとフラグ指定されていない残りの引数を取得できる
	log.Println(flag.Args())
}
