package main

import (
	"bufio"
	"fmt"
	"os"
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
	dup := make(map[string]struct{})
	for _, w := range words {
		dup[w] = struct{}{}
	}
	return len(dup) == len(words)

}
