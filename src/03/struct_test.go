package study03

import "testing"

func TestMemberAccess1(t *testing.T) {
	type Person struct {
		name string
	}
	p := &Person{
		name: "yamada",
	}
	// pは本来ポインタなのでデリファレンスが必要なはずだがメンバにアクセスする際は省略できる
	got1 := p.name    // このように書くこともできるが...
	got2 := (*p).name // 本来はこのように書くのが正しい
	expected := "yamada"
	if got1 != expected {
		t.Errorf("got1: %v, actual:%v", got1, expected)
	}
	if got2 != expected {
		t.Errorf("got2: %v, actual:%v", got2, expected)
	}
}

func TestMemberAccess2(t *testing.T) {
	name := "yamada"
	type Person struct {
		name *string
	}
	p := Person{
		name: &name,
	}
	// 先程の例とは違いメンバがポインタの場合は省略できない
	got := *(p.name)
	// このような書き方は許されない
	// got := p.name
	expected := "yamada"
	if got != expected {
		t.Errorf("got: %v, actual:%v", got, expected)
	}
}

func TestCreateStruct(t *testing.T) {
	type Person struct {
		name string
	}
	p1 := new(Person) // newで作成する方法
	var p2 Person     // var宣言で作成する方法
	p3 := &Person{    // リテラルで作成する方法
		name: "Yamada",
	}
	p1.name = "Yamada"
	p2.name = "Yamada"
	if !(p1.name == p2.name && p2.name == p3.name) {
		t.Errorf("name must be Yamada")
	}
}
