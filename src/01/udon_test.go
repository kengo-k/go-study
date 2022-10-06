// うどんのトッピング注文の実装例
package study01

import (
	"testing"
)

type Portion int

// 麺の量を指定する定数
const (
	Regular Portion = iota
	Small
	Large
)

// うどん構造体
type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func New(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{p, false, 2}
}

func (u *Udon) SetPortion(men Portion) *Udon {
	u.men = men
	return u
}

func (u *Udon) SetAburaage() *Udon {
	u.aburaage = true
	return u
}

func (u *Udon) SetEbiten(ebiten uint) *Udon {
	u.ebiten = ebiten
	return u
}

type Topping func(u *Udon)

func SetSize(men Portion) Topping {
	return func(u *Udon) {
		u.men = men
	}
}

func SetAburaage() Topping {
	return func(u *Udon) {
		u.aburaage = true
	}
}

func SetEbiten(ebiten uint) Topping {
	return func(u *Udon) {
		u.ebiten = ebiten
	}
}

func NewUdon(toppings ...Topping) *Udon {
	u := &Udon{}
	for _, t := range toppings {
		t(u)
	}
	return u
}

func TestSimpleOrder(t *testing.T) {
	kake := New(Large, false, 0)
	expectedSize := Large
	if kake.men != expectedSize {
		t.Errorf("got `%v` , expected `%v`", kake.men, expectedSize)
	}
}

func TestVariationOrder(t *testing.T) {
	tempura := NewTempuraUdon(Small)
	expectedCountEbiCount := uint(2)
	if tempura.ebiten != expectedCountEbiCount {
		t.Errorf("got `%v` , expected `%v`", tempura.ebiten, expectedCountEbiCount)
	}
}

func TestFluentOrder(t *testing.T) {
	pu := (&Udon{}).
		SetPortion(Large).
		SetAburaage().
		SetEbiten(3)
	expectedMen := Large
	expectedAge := true
	expectedEbi := uint(3)
	u := *pu
	if u.men != expectedMen {
		t.Errorf("got `%v` , expected `%v`", u.men, expectedMen)
	}
	if u.aburaage != expectedAge {
		t.Errorf("got `%v` , expected `%v`", u.aburaage, expectedAge)
	}
	if u.ebiten != expectedEbi {
		t.Errorf("got `%v` , expected `%v`", u.ebiten, expectedEbi)
	}
}

func TestOptionalOrder(t *testing.T) {
	pu := NewUdon(
		SetSize(Large),
		SetAburaage(),
		SetEbiten(3))
	expectedMen := Large
	expectedAge := true
	expectedEbi := uint(3)
	u := *pu
	if u.men != expectedMen {
		t.Errorf("got `%v` , expected `%v`", u.men, expectedMen)
	}
	if u.aburaage != expectedAge {
		t.Errorf("got `%v` , expected `%v`", u.aburaage, expectedAge)
	}
	if u.ebiten != expectedEbi {
		t.Errorf("got `%v` , expected `%v`", u.ebiten, expectedEbi)
	}
}
