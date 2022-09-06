package study01

import (
  "testing"
  "errors"
)

// エラーの値を利用者側で比較用の定数として使いたいが
// この場合はconstではビルドエラーになってしまう
// 関数の戻り値をconstで取れないため
// この場合はvarで宣言するしかない
var mutableError error = errors.New("mutableError")

// 上記の問題を解決するために定数値をErrorインタフェースに適合させる
type customError int

// Errorインタフェースを満たすにはErrorメソッドを定義する
func (e customError) Error() string {
  return "customError"
}

// constでエラー定数を宣言したため変更不可能にすることができる
const ImmutableError customError = 0

func TestError(t *testing.T) {

  // → 値を変更しようと思えばできてしまう...
  mutableError = errors.New("changedError")

  // ImmutableErrorはErrorインタフェースを満たすためError型に代入できる
  var e error = ImmutableError
  expectedMessage := "customError"
  if (e.Error() != expectedMessage) {
    t.Errorf("got `%v` , expected `%v`", e.Error(), expectedMessage)
  }
  // APIの戻り値のerrorと定数Errorを比較できる
  if (e != ImmutableError) {
    t.Errorf("Error, e must be ImmutableError")
  }
}