package word

import "testing"

func TestPalindroom(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`ISPalindrome("detartrated" = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`ISPalindrome("detartrated" = false`)
	}
}

func TestNonePaildroom(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`ISPalindrome("detartrated" = false`)
	}
}
