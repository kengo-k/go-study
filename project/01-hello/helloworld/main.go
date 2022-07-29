// 最初にパッケージ名がくる
// mainパッケージは特別なパッケージで実行形式毎に必ず含まれていなければならない
// パッケージ名はすべて小文字であること
package main

import (
  "fmt";
  "strconv"
)

// mainパッケージのmain関数がエントリポイントになる
// パッケージレベルで定義された関数や変数については
// 大文字の場合はpublicになる(パッケージ外から参照できる)
// 小文字の場合はprivate
func main() {
  hello()
  num()
  pointer()
  array()
}

func hello() {
  fmt.Println("Hello World")
  // varで変数宣言 変数名の後に型名が来る
  var hello string = "Hello"
  // ただし右辺で型が決まるなら型名は省略してよい
  var world = "World"
  // 宣言と同時に代入する場合はさらにvarも省略可能
  // ただし関数内部でしか使えない
  space := ` ` // 文字列リテラルはバッククオートも使用可能
  fmt.Println(hello + space + world)
}

func num() {
  // 使える数値の型はint, uint, およびintN, uintN, flontN(Nはビット数。8,16,32,64が使える)
  // 型指定なしの場合はintもしくはfloat64になる
  num := 123
  // 型変換の際は型名を関数として使用する
  var num2 float64 = float64(num)
  // 数値を文字列に変換するにはstrconv.FormatFloatを使う
  // 第2引数は表現形式: 数値をそのまま文字列にしたい場合は'f'でよさそう
  // 第3引数は精度: 基本的には-1でよさそう
  // 第4引数はビット数: 基本的には64でよさそう
  numstr := strconv.FormatFloat(num2, 'f', -1, 64)
  fmt.Println(numstr)

  strnum := "123.456"
  // 数値文字列を数値に変換するにはstrconv.ParseFloatを使う
  // 戻り値として変換後の値とエラーを返す
  // 未使用の変数があるとビルドエラーになってしまうので必ずチェックする必要がある
  num3, err := strconv.ParseFloat(strnum, 64)
  // 成功時errはnil
  if (err == nil) {
    num3str := strconv.FormatFloat(num3, 'f', -1, 64)
    fmt.Println(num3str)
  }
}

func pointer() {
  a := 100
  b := 200
  pa := &a
  var pb *int = nil
  pb = &b
  fmt.Println(*pa + *pb)
}

func array() {
  // 固定長配列
  var nums [3]int = [3]int{1,2,3}
  // lenで長さを取得
  fmt.Println(len(nums))
  // インデックスアクセス
  fmt.Println(nums[1])

  // 可変長配列
  nums2 := []int{5,6,7}
  // appendで追加
  nums2 = append(nums2, 8)
  nums2 = append(nums2, 9)
  // 切り出し
  nums3 := nums2[2:4]
  fmt.Println(len(nums3))
  fmt.Println(nums3[0])
}
