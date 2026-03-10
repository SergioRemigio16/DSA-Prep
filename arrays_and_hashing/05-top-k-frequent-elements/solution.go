package main

func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	// Create hashmap freqMap to hold appearance frequency
	freqMap := make(map[int]int) // Space: O(n)
	// Loop through nums and fill up freqMap with frequency of each int seen
	for _, num := range nums { // Time: O(n)
		freqMap[num]++
	}
	// Create a 2d array called freqBuckets of size n + 1 where every index is the frequency
	// and the 1d array contains integers with said frequency
	freqBuckets := make([][]int, n+1) // Space: O(n)
	// Loop through map
	for key, value := range freqMap { // Time: O(n)
		// freqBuckets[value] append key
		freqBuckets[value] = append(freqBuckets[value], key) // Time: O(1)
	}

	// Create solutionArr
	solutionArr := []int{}
	// Loop through each array in freqBuckets starting from the end
	for i := len(freqBuckets) - 1; i >= 0; i-- { // Time: O(n)
		currArray := freqBuckets[i]
		// If curr array is non empty
		if len(currArray) != 0 {
			// Loop through each integer in array
			for _, num := range currArray {
				// If k > 0
				if k > 0 {
					// append currInteger to solutionArr
					solutionArr = append(solutionArr, num)
				}
				// k--
				k--
				// If k == 0
				if k == 0 {
					return solutionArr
				}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	solution := topKFrequent(nums, k)
	for _, num := range solution {
		println(num)
	}
}
