package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	for scanner.Scan() {
		if valid(scanner.Text()) {
			n++
		}
	}
	fmt.Println(n)
}

func valid(phrase string) bool {
	words := strings.Split(phrase, " ")
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			//			fmt.Println(sorted(words[i]), sorted(words[j]), sorted(words[i]) == sorted(words[j]))
			if sorted(words[i]) == sorted(words[j]) {
				return false
			}
		}
	}
	return true
}

func sorted(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
