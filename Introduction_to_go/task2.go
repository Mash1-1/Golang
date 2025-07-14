// Task 2
package main

import (
	"fmt"
	"strings"
)

func freq_count(s string) map[string]int {
	mp := make(map[string]int)

	prev := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			prev += string(s[i])
		}else {
			prev = strings.ToLower(prev)
			mp[prev] += 1
			prev = ""
		}
	}
	mp[prev] += 1
	return mp
}

func IsPalindrome(s string) bool {
	left := 0 
	right := len(s) - 1

	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	greeting := "Hello World world"
	fmt.Println(freq_count(greeting))

	t1 := "aba"
	t2 := "ab"
	fmt.Printf("Test 1: %v\nTest 2: %v\n", IsPalindrome(t1), IsPalindrome(t2)) 
}
