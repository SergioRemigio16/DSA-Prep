package main

// An array of prefixes and suffixes is created.
// Prefix - Product of all numbers to the left of an index in nums
// Suffix - Product of all numbers to the right of index in nums.
// ProductExceptSelf[i] is found by simply multiplying Prefix[i] * Suffix[i]
func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefixes := make([]int, n)
	prefixes[0] = 1
	suffixes := make([]int, n)
	suffixes[n-1] = 1
	for i := 1; i < n; i++ {
		prefixes[i] = prefixes[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		suffixes[i] = suffixes[i+1] * nums[i+1]
	}

	productExceptSelf := make([]int, n)
	for i := 0; i < n; i++ {
		productExceptSelf[i] = prefixes[i] * suffixes[i]
	}

	return productExceptSelf
}

func main() {
	nums := []int{1, 2, 4, 6}
	println(productExceptSelf(nums)[0])
	println(productExceptSelf(nums)[1])
	println(productExceptSelf(nums)[2])
	println(productExceptSelf(nums)[3])
}
