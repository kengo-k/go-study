package study04

import (
	"testing"
)

type Animal interface {
	Run() string
}

type Sleeper interface {
	Sleep() string
}

type Cat struct{}

func (cat *Cat) Run() string {
	return "cat run...!"
}

func (cat *Cat) Sleep() string {
	return "cat sleep...!"
}

func TestCast(t *testing.T) {
	type Status int
	var numStatus int = 404
	// numStatusはint型でありStatusではないので代入(=ビルド)できない
	//var _ Status = numStatus

	// 型名(変数)で明示的にキャストすれば代入できる
	var _ Status = Status(numStatus)

	// CatはそのままAnimalに代入可能
	var cat Animal = &Cat{}
	// catはAnimalインタフェースでありSleeperであるかどうかは静的に判断できないためキャストはできない
	//var sleeper Sleeper = cat
	// インタフェースを静的にキャストできない場合は型アサーションを使い実行時に判定する
	if sleeper, ok := cat.(Sleeper); ok {
		got := sleeper.Sleep()
		expected := "cat sleep...!"
		if got != expected {
			t.Errorf("got: %v, expected: %v", got, expected)
		}
	} else {
		t.Errorf("cat should be Sleeper")
	}

}
