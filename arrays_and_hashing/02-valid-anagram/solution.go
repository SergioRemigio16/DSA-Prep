package main

import "reflect"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_char_freq_map := make(map[rune]int)
	t_char_freq_map := make(map[rune]int)

	// For every character S
	// i: index
	// r: rune (Unicode character)
	// Time: O(n)
	// Space: O(1)
	for _, r := range s {
		// Simply increase the value by one because if the map doesn't contain the
		// key then it will be created and set to 0
		s_char_freq_map[r]++
	}

	//Time: O(n)
	// Space: O(1)
	for _, r := range t {
		t_char_freq_map[r]++
	}

	// Return true if the two maps equal each other
	// Time: O(n+m)
	return reflect.DeepEqual(s_char_freq_map, t_char_freq_map)
	// Time Complexity: O(n) + O(m) + O(n + m) = O(2n + 2m) = O(n+m)
	// Space Complexity: O(1)
}

func main() {
	println("Hello, World!")
}
