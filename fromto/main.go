package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.PrintIfNot("hi"))
	fmt.Print(piscine.SearchReplace("hélla", "é", "o")) // if you add this
}
