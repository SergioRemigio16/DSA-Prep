package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var lcp strings.Builder
	firstWord := strs[0]
	for i := 0; i < len(firstWord); i++ {
		for _, word := range strs {
			if len(word) == i {
				return lcp.String()
			} else {
				if firstWord[i] == word[i] {
				} else {
					return lcp.String()
				}
			}
		}
		lcp.WriteByte(firstWord[i])
	}

	return lcp.String()
}

func main() {
	input := []string{"dance", "dag", "danger", "damage"}
	println(longestCommonPrefix(input))
}
