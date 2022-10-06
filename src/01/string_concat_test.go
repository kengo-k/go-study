package study01

import (
	"strings"
	"testing"
)

func TestStringConcat(t *testing.T) {
	words := []string{"Hello", "Go", "World"}
	// 文字列を+で連結するのは性能的によくないためbuilderで連結処理を行う
	var builder strings.Builder
	for i, word := range words {
		if i != 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(word)
	}
	actual := builder.String()
	expected := "Hello Go World"
	if actual != expected {
		t.Errorf("got `%v` , expected `%v`", actual, expected)
	}
}
