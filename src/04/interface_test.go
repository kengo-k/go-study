package study04

import (
	"fmt"
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

	var anyValue any = 1
	// 型switch
	switch v := anyValue.(type) {
	case string:
		t.Errorf("fail")
	case int:
		fmt.Printf("v is %v", v)
	default:
		t.Errorf("fail")
	}

}

type Container struct {
	runner Runner
}

func (c *Container) Run() string {
	return c.runner.Run()
}

func (c *Container) AddFactory(create func() Runner) {
	c.runner = create()
}

type Runner interface {
	Run() string
}

type RunnerImpl struct{}

func (r RunnerImpl) Run() string {
	return "RunnerImpl!"
}

var container = &Container{}

// init関数は最初に一度だけ呼び出される
// この処理の中で条件に応じて適切なFactoryを設定することで処理を切り替える
func init() {
	fmt.Printf("init called...")
	container.AddFactory(func() Runner {
		return &RunnerImpl{}
	})
}

func TestSwitchStatic(t *testing.T) {
	got := container.Run()
	expected := "RunnerImpl!"
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
