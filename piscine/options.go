package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	// If any argument has its first option char == 'h', print help and exit
	for _, a := range args {
		if len(a) >= 2 && a[0] == '-' && a[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
	}

	var opts uint32
	for _, a := range args {
		if len(a) < 2 || a[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		for i := 1; i < len(a); i++ {
			ch := a[i]
			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			opts |= uint32(1) << uint(ch-'a')
		}
	}

	// print 4 groups of 8 bits from high to low
	for byteIdx := 3; byteIdx >= 0; byteIdx-- {
		for bit := 7; bit >= 0; bit-- {
			idx := byteIdx*8 + bit
			if opts&(uint32(1)<<uint(idx)) != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		if byteIdx != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
