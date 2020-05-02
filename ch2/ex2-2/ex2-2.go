/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	String string
	Number int
	Input  string
)

func main() {
	if len(os.Args[1:]) != 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex2-2:%v\n", err)
				os.Exit(1)
			}
			fmt.Println(t + 1.0)
		}
	} else {
		//f := bufio.NewReader(os.Stdin) //读取输入的内容

		//Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
		//strings.TrimRight(Input,"\n")
		//Input = Input[0, len(Input)-1]
		/*t, err := strconv.ParseFloat(Input[0:len(Input)-1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2-2:%v\n", err)
			os.Exit(1)
		}

		context := strings.Fields(Input)
		for _, elem := range context {
			fmt.Println(elem)
		}





	}
}*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {

		fmt.Printf("%d\t%s\n", n, line)

	}
}