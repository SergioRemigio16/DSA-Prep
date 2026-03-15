package main

func majorityElement(nums []int) int {
	n := len(nums)
	numsFreqMap := make(map[int]int)
	for i := 0; i < n; i++ {
		numsFreqMap[nums[i]]++
		if val, _ := numsFreqMap[nums[i]]; val > (n / 2) {
			return nums[i]
		}
	}
	return 0
}

func main() {
	nums := []int{5, 5, 1, 1, 1, 5, 5}
	println(majorityElement(nums))
	nums2 := []int{2, 2, 2}
	println(majorityElement(nums2))
}
