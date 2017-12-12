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
	garbage := false
	var ngarbage int
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch c {
		case '<':
			if !garbage {
				garbage = true
			} else {
				ngarbage++
			}
		case '>':
			if garbage {
				garbage = false
			} else {
				panic("should not get > outside garbage " + fmt.Sprintf("%c at %d", c, i))
			}
		case '!':
			i++
		default:
			if garbage {
				ngarbage++
			}
		}
	}
	fmt.Println(ngarbage)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
