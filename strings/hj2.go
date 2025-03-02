package strings

import (
	"bufio"
	"fmt"
	"os"
)

func hj2() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	c := scanner.Text()[0]
	c1 := byte(0)
	if c >= 'a' && c <= 'z' {
		c1 = c - 'a' + 'A'
	}
	if c >= 'A' && c <= 'Z' {
		c1 = c - 'A' + 'a'
	}
	count := 0
	for i := range s {
		if s[i] == c || s[i] == c1 {
			count += 1
		}
	}
	return count
}

func hj2_solution() {
	fmt.Println(hj2())
}
