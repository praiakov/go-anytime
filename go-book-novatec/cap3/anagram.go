package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s, s2 string) bool {
	if len(s) != len(s2) {
		return false
	}

	str1 := strings.Split(s, "")
	str2 := strings.Split(s2, "")
	sort.Strings(str1)
	sort.Strings(str2)

	st := strings.Join(str1, ",")
	st2 := strings.Join(str2, ",")

	return st == st2
}

func main() {
	fmt.Println("Is anagram? ", isAnagram("pedra", "perda"))
}
