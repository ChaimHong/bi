package bi

import (
	"strconv"
)

func Bit2Int(b int) int {
	str := strconv.Itoa(b)
	bs := []byte(str)
	val := 0

	for j, i := uint(0), len(bs)-1; i >= 0; i-- {
		switch string(bs[i]) {
		case "0":
		case "1":
			val += (1 << j)
		default:
			panic("param is err")
		}
		j++
	}

	return val
}

func Int2Bit(b int64) string {
	str := ""
	for i := uint(0); i < uint(63); i++ {
		if b&int64(1<<i) == int64(1<<i) {
			str = "1" + str
		} else {
			str = "0" + str
		}
	}

	return str
}
