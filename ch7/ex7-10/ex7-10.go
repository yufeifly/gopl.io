package main

import (
	"fmt"
	"os"
	"sort"
)

type stringSlice []byte

func (s stringSlice) Len () int {return len(s)}
func (s stringSlice) Swap (i,j int) {
	s[i], s[j] = s[j], s[i]
}
func (s stringSlice) Less(i,j int) bool {
	return s[i] < s[j]
}

func IsPalindrome (s sort.Interface) bool {
	for i,j:=0,s.Len()-1;i<j; i,j = i+1,j-1 {
		if !(!s.Less(i,j) && !s.Less(j,i)) {
			return false
		}
	}
	return true
}

func main() {
	ss := stringSlice([]byte("abccbaa"))
	fmt.Fprintln(os.Stdout, IsPalindrome(ss))
}
