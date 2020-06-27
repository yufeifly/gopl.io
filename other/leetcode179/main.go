package main

import (
	"fmt"
	"sort"
	"strconv"
)

type intSlice []int

func (s *intSlice) Len() int {
	return len(*s)
}

func (s *intSlice) Less(i, j int) bool {
	iStr := strconv.FormatInt(int64((*s)[i]), 10)
	jStr := strconv.FormatInt(int64((*s)[j]), 10)

	return compare(iStr, jStr)
}

func (s *intSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func compare(a, b string) bool {
	return a+b > b+a
}

func largestNumber(nums []int) string {
	arr := intSlice(nums)
	sort.Sort(&arr)
	nums = arr
	str := ""
	for _, val := range nums {
		str += strconv.Itoa(val)
	}
	return wipeExtraZeros(str)
}

func wipeExtraZeros(s string) string {
	ret := "0"
	for k, val := range s {
		if val != '0' {
			ret = s[k:]
			return ret
		}
	}
	return ret
}

func main() {
	arr := []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}
	fmt.Println(largestNumber(arr))

}
