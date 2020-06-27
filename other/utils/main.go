package main

import (
	"code.byted.org/gopkg/logs"
	"fmt"
	"strconv"
	"strings"
)

var bits int32 = 2

// 这里的参数s是xx.xx.xx.xx的格式，填充位数后转为int64类型
func Version2Int(s string) (v int64) {
	sector := strings.Split(s, ".")
	for key, val := range sector {
		vLen := int32(len(val))
		if vLen < bits {
			val = zeros(bits-int32(len(val))) + val
		} else if vLen != bits {
			logs.Error("bits error")
		}
		sector[key] = val
	}
	vStr := strings.Join(sector, "")
	v, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		logs.Info("ParseInt fail")
	}
	return
}

// 这里的s是xxxxxxxx的整形数，将其转为xx.xx.xx.xx的字符串，有0的话可以去除
func Version2Str(s int64) (v string) {
	vStr := strconv.FormatInt(s, 10)
	var vArr []string
	for i := len(vStr); i > 0; i = i - 2 {
		if i-2 < 0 {
			vArr = append(vArr, vStr[0:i])
		} else {
			vArr = append(vArr, vStr[i-2:i])
		}
	}
	for i, j := 0, len(vArr)-1; i < j; i, j = i+1, j-1 {
		vArr[i], vArr[j] = vArr[j], vArr[i]
	}
	for k, val := range vArr {
		vArr[k] = wipeRedundantZeros(val)
	}
	v = strings.Join(vArr, ".")
	return
}

func zeros(cnt int32) (s string) {
	for i := int32(0); i < cnt; i++ {
		s += "0"
	}
	return
}

func wipeRedundantZeros(s string) string {
	for k, v := range s {
		if v != '0' {
			return s[k:]
		}
	}
	return "0"
}

/*func allZeros(s string) bool {
	for _, v := range s {
		if v != '0' {
			return false
		}
	}
	return true
}*/

func main() {
	fmt.Println(Version2Int("1.0.1"))
	fmt.Println(Version2Str(10001))
}
