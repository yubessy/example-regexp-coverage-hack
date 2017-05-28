package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/yubessy/golang-regexp-customize/regexp"
)

func main() {
	re := regexp.MustCompile(`.+@.+\.(com|jp)`)
	fmt.Printf("%q\n", re.FindString("seafood"))
	fmt.Printf("%q\n", re.FindString("meat"))
	pp.Println(re.Inst())
}
