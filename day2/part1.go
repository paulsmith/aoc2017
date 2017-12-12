package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.Comma = '\t'
	r.FieldsPerRecord = -1

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("reading records: %v", err)
	}

	const maxint = int((^uint(0)) >> 1)
	const minint = -maxint - 1
	checksum := 0
	for _, row := range rows {
		largest, smallest := minint, maxint
		for i := range row {
			v := atoi(row[i])
			if v > largest {
				largest = v
			}
			if v < smallest {
				smallest = v
			}
		}
		checksum += largest - smallest
	}

	fmt.Printf("%d\n", checksum)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid int " + s)
	}
	return int(n)
}
