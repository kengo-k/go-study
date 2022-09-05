package study00

import (
  "testing"
  "fmt"
  "strconv"
  "os"
  "io"
  "encoding/json"
)

func TestHello(t *testing.T) {  
  // varで変数宣言 変数名の後に型名が来る
  var hello string = "Hello"
  // ただし右辺で型が決まるなら型名は省略してよい
  var world = "World"
  // 宣言と同時に代入する場合はさらにvarも省略可能
  // ただし関数内部でしか使えない
  space := ` ` // 文字列リテラルはバッククオートも使用可能
  
  strcat := hello + space + world
  expected := "Hello World"
  if (strcat != expected) {
    t.Errorf("got `%v` , expected `%v`", strcat, expected)
  }
}

func TestNum(t *testing.T) {
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
  expectedNumStr := "123"
  if (numstr != expectedNumStr) {
    t.Errorf("got `%v` , expected `%v`", numstr, expectedNumStr)
  }

  strnum := "123.456"
  // 数値文字列を数値に変換するにはstrconv.ParseFloatを使う
  // 戻り値として変換後の値とエラーを返す
  // 未使用の変数があるとビルドエラーになってしまうので必ずチェックする必要がある
  parsedFloat, err := strconv.ParseFloat(strnum, 64)
  // 成功時errはnil
  if (err != nil) {
    //num3str := strconv.FormatFloat(num3, 'f', -1, 64)
    t.Errorf("parse error, input=`%v`", strnum)
  }

  expectedParsedFloat := 123.456
  if (parsedFloat != expectedParsedFloat) {
    t.Errorf("got `%v` , expected `%v`", parsedFloat, expectedParsedFloat)
  }
}

func TestPointer(t *testing.T) {
  a := 100
  b := 200
  // 変数aのポインタをpaに格納
  pa := &a
  expectedPa := 100
  // *でポインタが指す値を取得する
  if (*pa != expectedPa) {
    t.Errorf("got `%v` , expected `%v`", *pa, expectedPa)
  }

  // ポインタ型にはnilを設定できる
  var pb *int = nil
  pb = &b

  sum := *pa + *pb
  expectedSum := 300
  if (sum != expectedSum) {
    t.Errorf("got `%v` , expected `%v`", sum, expectedSum)
  }
}
