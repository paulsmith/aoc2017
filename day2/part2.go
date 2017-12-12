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

	checksum := 0
	for _, row := range rows {
	row:
		for i := range row {
			v := atoi(row[i])
			for j := range row {
				if i == j {
					continue
				}
				u := atoi(row[j])
				if v%u == 0 {
					checksum += v / u
					break row
				} else if u%v == 0 {
					checksum += u / v
					break row
				}
			}
		}
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
