package study03

import "testing"

func TestMemberAccess1(t *testing.T) {
	type Person struct {
		name string
	}
	p := &Person{
		name: "yamada",
	}
	// TODO comment
	got1 := p.name
	got2 := (*p).name
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
	// TODO comment
	got := *(p.name)
	// NG
	// got := p.name
	expected := "yamada"
	if got != expected {
		t.Errorf("got: %v, actual:%v", got, expected)
	}
}
