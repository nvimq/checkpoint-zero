package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	input := strings.TrimSpace(os.Args[1])
	if input == "" {
		fmt.Println("Error")
		return
	}

	tokens := strings.Fields(input)
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "%":
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				result = a / b
			case "%":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				result = a % b
			}
			stack = append(stack, result)
		default:
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Println("Error")
				return
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}
