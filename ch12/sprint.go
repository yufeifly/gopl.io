package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

type Hello struct {
	s string
}

func (h *Hello) String() string {
	return "hello world"
}

func main() {
	h := Hello{
		s: "he",
	}
	/*fmt.Println(Sprint(1))
	fmt.Println(Sprint("abc"))
	fmt.Println(Sprint(true))
	fmt.Println(Sprint(&h))*/
	t := reflect.TypeOf(&h)
	fmt.Println(t.String())

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
}
