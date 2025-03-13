package stack

import (
	"path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}
	pop := func() {
		stack = stack[0 : len(stack)-1]
	}
	push := func(c string) {
		stack = append(stack, c)
	}
	push("/")
	paths := strings.Split(path, "/")
	for _, p := range paths {
		switch p {
		case "..":
			// 要处理开始就是/..情况
			if len(stack) > 1 {
				pop()
			}
		case ".", "":
			continue
		default:
			push(p)
		}
	}
	return filepath.Join(stack...)
}
