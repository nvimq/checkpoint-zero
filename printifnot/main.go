package main

import (
	"fmt"

	"github.com/nvimq/checkpoint-zero/piscine"
)

func main() {
	fmt.Print(piscine.PrintIfNot("abcdefz"))
	fmt.Print(piscine.PrintIfNot("abc"))
	fmt.Print(piscine.PrintIfNot(""))
	fmt.Print(piscine.PrintIfNot("14"))
}
