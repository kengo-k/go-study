package study01

import (
  "testing"
)

type status int
type flag int

// 1始まりの連番が定義される
const (
  SUCCESS status = iota + 1
  FAIL
  WARNING
)

// シフトして値を2倍にしていく定数
const (
  READ flag = 1 << iota
  WRITE
)

func TestSimpleIota(t *testing.T) {
  var st status = FAIL
  expectedSt := FAIL
  if (st != expectedSt) {
    t.Errorf("got `%v` , expected `%v`", st, expectedSt)
  }

  // status型はあくまでもint(の別名)なので想定していた定数の範囲外の値を入れてもビルドできてしまう
  var st2 status = 100
  var expectedSt2 status = 100
  // expectedSt2 := 100
  // ただし↑のようにした場合はexpectedSt2はintになるため、intとstatusが型が違うため直接比較することができずにビルドエラーとなる
  if (st2 != expectedSt2) {
    t.Errorf("got `%v` , expected `%v`", st2, expectedSt2)
  }
}

func TestFlagIota(t *testing.T) {
  if (READ != 1) {
    t.Errorf("got `%v` , expected `%v`", READ, 1)
  }
  if (WRITE != 2) {
    t.Errorf("got `%v` , expected `%v`", WRITE, 2)
  }
  if (READ|WRITE != 3) {
    t.Errorf("got `%v` , expected `%v`", READ|WRITE, 3)
  }
}