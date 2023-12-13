package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("app/day10/sample.csv")
	if err != nil {
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		fmt.Println(record)
	}
}
