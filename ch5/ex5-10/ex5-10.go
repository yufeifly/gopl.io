package main

import (
	"fmt"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
func main() {
	for k, v := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", k, v)
	}
}

func topoSort(m map[string][]string) map[int]string {
	var order = make(map[int]string)
	index := 1
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[index] = item
				index++
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	visitAll(keys)
	return order
}
