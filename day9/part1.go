package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(b)
	var level int
	garbage := false
	var score int
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch c {
		case '{':
			if !garbage {
				level++
			}
		case '}':
			if !garbage {
				score += level
				level--
			}
		case '<':
			if !garbage {
				garbage = true
			}
		case '>':
			if garbage {
				garbage = false
			} else {
				panic("should not get > outside garbage " + fmt.Sprintf("%c at %d", c, i))
			}
		case '!':
			i++
		}
	}
	fmt.Println(score)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
