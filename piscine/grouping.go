package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	pat := os.Args[1]
	txt := os.Args[2]

	if len(pat) < 2 || pat[0] != '(' || pat[len(pat)-1] != ')' {
		return
	}
	inside := pat[1 : len(pat)-1]
	if inside == "" {
		return
	}

	// split pattern by '|', invalid if any part empty
	parts := []string{}
	cur := ""
	for _, r := range inside {
		if r == '|' {
			if cur == "" {
				return
			}
			parts = append(parts, cur)
			cur = ""
		} else {
			cur += string(r)
		}
	}
	if cur == "" {
		return
	}
	parts = append(parts, cur)

	// split text into words by whitespace, then trim punctuation
	words := []string{}
	cur = ""
	for _, r := range txt {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			if cur != "" {
				if t := trim(cur); t != "" {
					words = append(words, t)
				}
				cur = ""
			}
		} else {
			cur += string(r)
		}
	}
	if cur != "" {
		if t := trim(cur); t != "" {
			words = append(words, t)
		}
	}

	// for each word, for each pattern part, print once if the part is present
	count := 0
	for _, w := range words {
		for _, p := range parts {
			if contains(w, p) {
				count++
				fmt.Printf("%d: %s\n", count, w)
			}
		}
	}
}

func isWordRune(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}
	if r == '\'' || r == '\u2019' {
		return true
	}
	return false
}

func trim(s string) string {
	rs := []rune(s)
	i := 0
	for i < len(rs) && !isWordRune(rs[i]) {
		i++
	}
	j := len(rs) - 1
	for j >= i && !isWordRune(rs[j]) {
		j--
	}
	if i > j {
		return ""
	}
	return string(rs[i : j+1])
}

func contains(s, sub string) bool {
	rs := []rune(s)
	rp := []rune(sub)
	if len(rp) == 0 {
		return true
	}
	for i := 0; i+len(rp) <= len(rs); i++ {
		ok := true
		for j := range rp {
			if rs[i+j] != rp[j] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
