package main

import (
	"fmt"

	"github.com/yubessy/regexp-coverage-hack/regexp"
)

func main() {
	re := regexp.MustCompile(`[a-z]+\.(co|ne)\.jp`)
	re.MatchString("hoge.co.jp")
	c, a := re.Coverage()
	fmt.Printf("%d / %d", c, a)
}
