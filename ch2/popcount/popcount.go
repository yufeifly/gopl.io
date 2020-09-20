package popcount

//import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

// 使用的是动态规划的办法，把最后一位和其他位分开，pc[i/2]表示其他位的1的个数，byte(i&1)表示最后一位的1的个数
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("%c ",pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	/*for _, val := range pc {
		fmt.Println(val)
	}*/

	// 使用循环
	/*var answer int
	for i:=0;i<8;i++ {
		answer += int(pc[byte(x>>(i*8))])
	}
	return answer*/

	// 使用每次移动1bit
	/*var answer int
	for x > 0 {
		if x & 1 == 1 {
			answer += 1
		}
		x = x >> 1
	}
	return answer*/

	// 使用x&(x-1)
	/*	var answer int
		for x > 0 {
			answer += 1
			x = x & (x-1)
		}*/
	//return answer

	// 使用单一表达式
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

}

func PopCount2(x uint64) int {
	/*for _, val := range pc {
		fmt.Println(val)
	}*/

	// 使用循环
	/*var answer int
	for i:=0;i<8;i++ {
		answer += int(pc[byte(x>>(i*8))])
	}
	return answer*/

	// 使用每次移动1bit
	/*var answer int
	for x > 0 {
		if x & 1 == 1 {
			answer += 1
		}
		x = x >> 1
	}
	return answer*/

	// 使用x&(x-1)
	/*	var answer int
		for x > 0 {
			answer += 1
			x = x & (x-1)
		}*/
	//return answer

	// 使用单一表达式
	ans := 0
	for i := 0; i < 8; i++ {
		ans += int(pc[byte(x>>(i*8))])
	}
	return ans
	/*return int(pc[byte(x>>(0*8))] +
	pc[byte(x>>(1*8))] +
	pc[byte(x>>(2*8))] +
	pc[byte(x>>(3*8))] +
	pc[byte(x>>(4*8))] +
	pc[byte(x>>(5*8))] +
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))])*/

}

func PopCount3(x uint64) int {
	/*for _, val := range pc {
		fmt.Println(val)
	}*/

	// 使用循环
	/*var answer int
	for i:=0;i<8;i++ {
		answer += int(pc[byte(x>>(i*8))])
	}
	return answer*/

	// 使用每次移动1bit
	var answer int
	for x > 0 {
		if x&1 == 1 {
			answer += 1
		}
		x = x >> 1
	}
	return answer

	// 使用x&(x-1)
	/*	var answer int
		for x > 0 {
			answer += 1
			x = x & (x-1)
		}*/
	//return answer

	// 使用单一表达式
	/*ans := 0
	for i:=0;i<8;i++ {
		ans += int(pc[byte(x>>(i*8))])
	}
	return ans*/

}

func PopCount4(x uint64) int {
	/*for _, val := range pc {
		fmt.Println(val)
	}*/

	// 使用循环
	/*var answer int
	for i:=0;i<8;i++ {
		answer += int(pc[byte(x>>(i*8))])
	}
	return answer*/

	// 使用x&(x-1)
	var answer int
	for x > 0 {
		answer += 1
		x = x & (x - 1)
	}
	return answer

	// 使用单一表达式
	/*ans := 0
	for i:=0;i<8;i++ {
		ans += int(pc[byte(x>>(i*8))])
	}
	return ans*/

}
