package csv

import (
	"encoding/csv"
	"os"
)

func Read(file string) (contents []string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	c := csv.NewReader(f)
	strs, err := c.Read()
	if err != nil {
		panic(err)
	}
	return strs
}

func ReadAll(file string) (lines [][]string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	c := csv.NewReader(f)
	lines, err = c.ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}
