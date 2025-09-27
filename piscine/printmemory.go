package piscine

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	hex := []rune("0123456789abcdef")

	for i, b := range arr {
		// Печатаем hex
		z01.PrintRune(hex[b>>4])
		z01.PrintRune(hex[b&0x0F])

		// Определяем — перенос строки или пробел
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 { // пробел только если это не последний байт
			z01.PrintRune(' ')
		}
	}

	// Перенос строки, если не закончили на кратном 4-х байте
	if len(arr)%4 != 0 {
		z01.PrintRune('\n')
	}

	// ASCII-часть
	for _, b := range arr {
		r := rune(b)
		if unicode.IsGraphic(r) {
			z01.PrintRune(r)
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
