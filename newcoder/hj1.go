package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hj1() int {
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		return 0
	}

	parts := strings.Fields(result)
	if len(parts) <= 0 {
		return 0
	}

	return len(parts[len(parts)-1])
}

func hj1_solution() {
	fmt.Print(hj1())
}
