package main

import "fmt"

func caesarCipher(s string, k int32) string {
	k = k % 26
	res := make([]byte, len(s))
	var end int32 = 0
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			end = 'Z'
		} else if c >= 'a' && c <= 'z' {
			end = 'z'
		} else {
			res[i] = byte(c)
			continue
		}

		if c+k > end {
			res[i] = byte(c + k - 26)
		} else {
			res[i] = byte(c + k)
		}
	}
	return string(res)
}

func camelcase(s string) int32 {
	var res int32 = 1
	for j := 1; j < len(s); j++ {
		if rune(s[j]) >= rune('A') && rune(s[j]) <= rune('Z') {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(caesarCipher("Always-Look-on-the-Bright-Side-of-Life", 5))
	fmt.Println(camelcase("saveChangesInTheEditor"))
}
