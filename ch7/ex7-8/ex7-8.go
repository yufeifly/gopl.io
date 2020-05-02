package main

import (
	"fmt"
	"sort"
)

type book struct {
	name   string
	price  float64
	author string
}
type byFunc func(i, j int) bool
type tableSlice struct {
	lists     []*book
	opInd 	  []int
	lessFuncs []byFunc
}

func (ts tableSlice) Len() int {
	return len(ts.lists)
}
func (ts tableSlice) Swap(i, j int) {
	ts.lists[i], ts.lists[j] = ts.lists[j], ts.lists[i]
}
func (ts tableSlice) Less(i, j int) bool {
	for _, opI := range ts.opInd {
		if ts.lessFuncs[opI](i, j) {
			return true
		} else if !ts.lessFuncs[opI](j, i) {
			continue
		} else {
			return false
		}
	}
	/*for t := len(ts.lessFuncs) - 1; t >= 0; t-- {

	}
	*/
	return false
}

func (ts tableSlice) byName(i, j int) bool {
	return ts.lists[i].name < ts.lists[j].name
}
func (ts tableSlice) byPrice(i, j int) bool {
	return ts.lists[i].price < ts.lists[j].price
}

func (ts tableSlice) byAuthor(i, j int) bool {
	return ts.lists[i].author < ts.lists[j].author
}

func (ts *tableSlice) updateOpInd(i int) {
	OpIndLen := len(ts.opInd)
	if i >= OpIndLen {
		return
	}
	iVal := ts.opInd[i]
	for j:=i-1; j>=0;j-- {
		ts.opInd[j+1] = ts.opInd[j]
	}
	ts.opInd[0] = iVal
}

func main() {
	book1 := book{"GoLang", 65.50, "Aideng"}
	book2 := book{"GoLang", 45.50, "yufei"}
	book3 := book{"PHP", 45.50, "Sombody"}
	book4 := book{"PHP", 13.50, "yufei"}
	book5 := book{"Z", 45.50, "Tan"}
	ts := tableSlice{
		lists: []*book{&book1, &book2, &book3, &book4, &book5},
	}
	ts.lessFuncs = []byFunc{ts.byName, ts.byPrice,ts.byAuthor}

	for i, _ := range ts.lessFuncs {
		ts.opInd = append(ts.opInd,i)
	}

	sort.Sort(ts)
	for _, book := range ts.lists {
		fmt.Println(*book)
	}
	fmt.Println(ts.opInd)

	fmt.Println("----------")

	ts.updateOpInd(1)
	sort.Sort(ts)
	for _, book := range ts.lists {
		fmt.Println(*book)
	}
	fmt.Println(ts.opInd)

	fmt.Println("----------")
	ts.updateOpInd(2)
	sort.Sort(ts)
	for _, book := range ts.lists {
		fmt.Println(*book)
	}
	fmt.Println(ts.opInd)

}