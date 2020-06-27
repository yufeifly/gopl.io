package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	numsLen := len(nums)
	if numsLen <= 1 {
		return false
	}
	mp := map[int]int{nums[0]: 0}
	for i := 1; i < numsLen; i++ {
		if _, ok := mp[nums[i]]; ok {
			if i-mp[nums[i]] <= k {
				return true
			}
		}
		mp[nums[i]] = i
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(arr, 3))
}
