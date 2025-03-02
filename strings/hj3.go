package strings

import (
	"fmt"
)

func hj3() {
	bitmap := [8]int64{0}
	var count int
	fmt.Scan(&count)
	var num int
	for i := 0; i < count; i += 1 {
		fmt.Scan(&num)
		num -= 1 // 从0为开始保存
		bitmap[num/64] |= (1 << (num % 64))
	}
	for idx, num := range bitmap {
		for j := 0; j < 64; j++ {
			if num&1 != 0 {
				fmt.Println(idx*64 + j + 1)
			}
			num >>= 1
		}
	}
}
