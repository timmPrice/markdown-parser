package main

import (
	"fmt"
	"markdown-parser/src"
)

func main() {
	fmt.Println(src.LookupToken("*"))
	fmt.Println(src.LookupToken("##"))
	fmt.Println(src.LookupToken("abc"))
}
