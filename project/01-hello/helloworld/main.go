// 最初にパッケージ名がくる
// mainパッケージは特別なパッケージで実行形式毎に必ず含まれていなければならない
// パッケージ名はすべて小文字であること
package main

import (
  "fmt"
  "strconv"
  "os"
  "io"
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
  maptest()
  iftest()
  fortest()
  switchtest()
  functest()
  defertest()
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

func maptest() {
  // mapを作成するには`map`か`make`を使用する
  status := map[int]string {
    200: "OK",
    404: "Not Found",
  }
  // mapアクセスは二番目の戻り値として見つかったかを返すが取らないことも可能
  fmt.Println(status[200])

  persons := make(map[string]int)
  persons["Yamada"] = 42
  persons["Tanaka"] = 51
  // 二番目の戻り値を取って存在チェック
  yamada, exists := persons["Yamada"]
  if (exists) {
    fmt.Println(yamada)
  }
}

func iftest() {
  num := 1
  b := true
  // 他の言語にあるようないい感じにbool値として判断してくれる機能はないためbool値を明示する必要がある
  if (num == 0) {
    fmt.Println(num)
    // ifの後のかっこはなくてもいい
  } else if b {
    fmt.Println(b)
  } else {
    fmt.Println(b)
  }

  status := map[int]string {
    200: "OK",
    404: "Not Found",
  }

  // 結果の取得と存在チェックをひとつのifで記述できる
  result := 404
  if statusText, exists := status[result]; exists {
    fmt.Println(statusText)
  }
}

func fortest() {
  names := []string{"Yamada", "Tanaka", "Suzuki"}
  // rangeは
  // 配列の場合index, valueを返す
  // mapの場合はkey, valueを返す
  // 二番目がいらない場合は省略してもよい
  // 二番目だけがほしい場合はアンダースコアを使って使用しないことを明示する
  for i, s := range names {
    fmt.Printf("[%v]: %v\n", i, s)
  }
  for _, s := range names {
    fmt.Printf("%v\n", s)
  }

  // 他の言語と同様にbreakやcontinueが使える
  nums := []int{1,2,3,4,5}
  for _, n := range nums {
    if n >3 {
      break
    }
    fmt.Printf("value is %v\n", n)
  }

  // いわゆるwhile
  i := 0
  for i < len(nums) {
    fmt.Printf("nums[%v]=%v\n", i, nums[i])
    i++
  }

  i = 0
  // whileの無限ループ
  for {
    if i > len(nums) {
      break
    }
    i++
  }
  fmt.Println("breaked")

  // 伝統的なfor
  for i := 0; i < len(nums); i++ {
    fmt.Printf("[%v]: %v\n", i, nums[i])
  }
}

func switchtest() {
  status := 404
  switch status {
  case 200:
    // 他の言語のswitchのようにデフォルトではfallthroughしないため
    // breakする必要がない
    fmt.Println("OK")
  case 404:
    fmt.Println("Not Found")
    // fallthroughしたい場合
    fallthrough
  case 500:
    fmt.Println("Server Error")
    fallthrough
  default:
    fmt.Println("Default")
  }
}

func functest() {
  a := 5
  b := 12
  fmt.Printf("%v + %v = %v\n", a, b, add(a, b))

  _, err := div(7, 0)
  if (err) {
    fmt.Println("ERROR!!")
  }
  ret, _ := div(7, 2)
  fmt.Printf("7 / 2 = %v\n", ret)

  add100 := addFactory(100)
  fmt.Printf("add100(5) = %v\n", add100(5))
}

// 一般的な関数定義
func add(a int, b int) int {
  return a + b
}

// a, bが同じ型なので型をまとめて定義している
func div(a, b int) (ret int, err bool) {
  // 戻り値の変数名を定義することができる
  // 既存の言語のようにreturnで値を返すのではなく変数に値を設定して単にreturnとだけすればよい

  if b == 0 {
    err = true
    return
  }
  // int同士で割り算した場合は結果はintとなる
  ret = a / b
  err = false
  return
}

// 関数を返す関数と無名関数
func addFactory(a int) func(int) int {
  f := func(b int) int {
    return a + b
  }
  return f
}

func defertest() {
  f, err := os.Create("sample.txt")
  if err != nil {
    fmt.Println("err", err)
    return
  }
  // deferを使うと後処理を定義できる
  // この関数を抜けたあとにファイルをクローズする
  defer f.Close()
  // すでにCloseを実行しているのでエラーになってしまうように見えてしまうが
  // Closeはこの関数が終了した後に行われるためファイル出力は正常に行われる
  io.WriteString(f, "Hello, Go")
}
