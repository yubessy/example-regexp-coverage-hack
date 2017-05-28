package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/yubessy/golang-regexp-customize/regexp"
)

func main() {
	re := regexp.MustCompile(`.+\.(co|ne)\.jp`)
	fmt.Printf("%q\n", re.FindString("hoge.co.jp"))
	fmt.Printf("%q\n", re.FindString("fuga.ne.jp"))
	pp.Println(re.Inst())
}
