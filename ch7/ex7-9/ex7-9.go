package main

import (
	"html/template"
	"os"
)

type book struct {
	Name   string
	Price  float64
	Author string
}
type byFunc func(i, j int) bool
type tableSlice struct {
	Lists     []*book
	opInd 	  []int
	lessFuncs []byFunc
}

func (ts tableSlice) Len() int {
	return len(ts.Lists)
}
func (ts tableSlice) Swap(i, j int) {
	ts.Lists[i], ts.Lists[j] = ts.Lists[j], ts.Lists[i]
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
	return false
}

func (ts tableSlice) byName(i, j int) bool {
	return ts.Lists[i].Name < ts.Lists[j].Name
}
func (ts tableSlice) byPrice(i, j int) bool {
	return ts.Lists[i].Price < ts.Lists[j].Price
}

func (ts tableSlice) byAuthor(i, j int) bool {
	return ts.Lists[i].Author < ts.Lists[j].Author
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

var issueList = template.Must(template.New("issuelist").Parse(`
<table>
<tr style='text-align: left'>
  <th>Name</th>
  <th>Price</th>
  <th>Author</th>
</tr>
{{range .Lists}}
<tr>
  <td>{{.Name}}</td>
  <td>{{.Price}}</td>
  <td>{{.Author}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	book1 := book{"GoLang", 65.50, "Aideng"}
	book2 := book{"GoLang", 45.50, "yufei"}
	book3 := book{"PHP", 45.50, "Sombody"}
	book4 := book{"PHP", 13.50, "yufei"}
	book5 := book{"Z", 45.50, "Tan"}
	ts := tableSlice{
		Lists: []*book{&book1, &book2, &book3, &book4, &book5},
	}
	ts.lessFuncs = []byFunc{ts.byName, ts.byPrice,ts.byAuthor}

	for i, _ := range ts.lessFuncs {
		ts.opInd = append(ts.opInd,i)
	}

	/*sort.Sort(ts)
	for _, book := range ts.Lists {
		fmt.Println(*book)
	}
	fmt.Println(ts.opInd)

	fmt.Println("----------")

	ts.updateOpInd(1)
	sort.Sort(ts)
	for _, book := range ts.Lists {
		fmt.Println(*book)
	}
	fmt.Println(ts.opInd)

	fmt.Println("----------")
	ts.updateOpInd(2)
	sort.Sort(ts)
	for _, book := range ts.Lists {
		fmt.Println(*book)
	}
	fmt.Println(ts.opInd)

	 */

	err := issueList.Execute(os.Stdout, ts)
	if err != nil { panic(err) }

}