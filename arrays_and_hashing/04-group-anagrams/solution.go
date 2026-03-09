package main

// n = number of words
// s = word length
// Time Complexity: O(n) * (O(s) + O(1)) + O(n) *O(s) = 2O(ns) = O(ns)
// Space Complexity: O(n) * O(1) + O(n) = 2O(n) = O(n)
func groupAnagrams(strs []string) [][]string {
	// Every key contains all words that are anagrams
	// Space : O(1)
	anagramMap := make(map[[26]int][]string)
	//go through every string in the array strings
	// Time: O(n)
	// Space : O(n)
	for _, word := range strs {
		// turn each string into a string array of size 26 with the frequency of each letter
		// Space : O(1)
		var arr [26]int
		// For every letter in the word
		// Time: O(s)
		// Space : O(1)
		for _, rune := range word {
			// increase the number of times this letter has appeared
			arr[rune-97]++
		}
		// Time: O(1)
		anagramMap[arr] = append(anagramMap[arr], word)
	}
	solutionArray := [][]string{}
	// Time: O(n)
	// Space: O(n)
	for _, value := range anagramMap {
		// Time: O(s)
		// Space: O(1)
		solutionArray = append(solutionArray, value)
	}
	return solutionArray
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	solution := groupAnagrams(strs)
	for _, anagrams := range solution {
		for _, word := range anagrams {
			print(word + " ")
		}
		println()
	}
	testStr := "hello World"
	println(testStr[0])
	println('a' - 97)
	println('z' - 97)

}
