package main

func groupAnagrams(strings []string) [][]string {
	//go through every string in the array strings
	for i, word := range strings {
		// turn each string into a string array of size 27 with the frequency of each letter
		var arr [27]int
		arr[0] = i
		for j, rune := range word {
			arr[j] = rune
		}
	}
	// The first index for the index position of the word in the strings array and 26 indices for the letters
	// add that array into a hashmap
	// iterate through hashmap and fill up [][]string that needs to be returned

	return nil
}

func main() {
}
