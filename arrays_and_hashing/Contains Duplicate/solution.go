package main

func hasDuplicate(nums []int) bool {
	var seen map[int]struct{} = make(map[int]struct{})
	// 
	for _, n := range nums {
		if _, exists := seen[n]; exists {
			return true
		}
		seen[n] = struct{}{}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2}
	hasDuplicate(nums)
	println(hasDuplicate(nums))
}
