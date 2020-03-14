package main

import (
    "fmt"
    "os"
)
// Check whether the characters in s can be replaced to get t
func can_map(s_str, t_str string) bool {
    // empty string and length mismatch always return false
	if s_str == "" || t_str == "" || len(s_str) != len(t_str) {
		return false
	}
	s := []rune(s_str)
	t := []rune(t_str)
	mapping := make(map[rune]rune)
    for i := 0; i < len(s); i++ {
		if val, ok := mapping[s[i]]; ok {
			// false if mapping conflicts
			if val != t[i] {
				return false
			}
		} else {
			// else create new mapping to char s[i]
			mapping[s[i]] = t[i]
		}
    }
    return  true
}


func main() {
	if len(os.Args) < 3 {
		fmt.Printf("false")
	} else {
		s := os.Args[1]
		t := os.Args[2]
		if can_map(s, t) {
			fmt.Printf("true")
		} else {
			fmt.Printf("false")
		}
	}
}