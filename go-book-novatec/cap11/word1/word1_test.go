package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartreated") {
		t.Error(`IsPalindrome("detartreated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Errorf(`IsPalindrome("palindrome") = true`)
	}
}
