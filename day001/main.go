package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	elfes := strings.Split(input, "\n\n")

	var max int
	var cals []int

	for i := range elfes {
		rows := strings.Split(elfes[i], "\n")

		sum := 0
		for j := range rows {
			c, _ := strconv.Atoi(rows[j])
			sum += c
		}

		if sum > max {
			max = sum
		}

		cals = append(cals, sum)
	}

	println(max)

	sort.Slice(cals, func(i, j int) bool {
		return cals[i] > cals[j]
	})

	println(cals[0] + cals[1] + cals[2])
}
