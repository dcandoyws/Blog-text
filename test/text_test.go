package test

import (
	"testing"
	"unicode"
)

func IsReadaer(is string) bool {
	var str []rune
	for _, r := range is {
		if unicode.IsLetter(r) {
			str = append(letters, unicode.ToLower(r))
		}
	}
	for i := range str {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func BMIsReadaer(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		IsReadaer("hello,this is testing!")
	}
}
