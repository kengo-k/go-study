package study03

import (
	"reflect"
	"testing"
)

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

func TestEmbedStruct(t *testing.T) {
	type Author struct {
		name string
	}
	// Book構造体にAuthor構造体を埋め込むことができる
	type Book struct {
		title string
		Author
	}
	book := &Book{
		title: "title1",
		Author: Author{
			name: "author1",
		},
	}
	expected := "author1"
	// 本来はbook.Author.nameだがショートカットbook.nameと書くことができる
	if book.name != expected {
		t.Errorf("got: %v, expected: %v", book.name, expected)
	}
	if book.Author.name != expected {
		t.Errorf("got: %v, expected: %v", book.Author.name, expected)
	}
}

func TestStructTag(t *testing.T) {
	type Person struct {
		Name string `map:"str"`
		Age  int    `map:"int"`
	}
	mapSrc := map[string]string{
		"name": "Yamada",
		"age":  "32",
	}
	var p Person
	_decode := func(e reflect.Value, src map[string]string) error {
		// TODO 後で書く
		return nil
	}
	decode := func(dest interface{}, src map[string]string) error {
		v := reflect.ValueOf(dest)
		e := v.Elem()
		return _decode(e, src)
	}
	decode(&p, mapSrc)
}
