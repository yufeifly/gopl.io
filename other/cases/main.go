package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "heh123HdsfHHHOOI-inhouse"
	sLow := strings.ToLower(s)
	sBool := strings.Contains(sLow, "inhouse")
	fmt.Println(sBool)
}
