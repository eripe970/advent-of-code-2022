package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type cube [3]float32

var offsets = []cube{{0.5, 0, 0}, {-0.5, 0, 0}, {0, 0.5, 0}, {0, -0.5, 0}, {0, 0, 0.5}, {0, 0, -0.5}}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sides := make(map[cube]int)

	for _, line := range lines {
		splits := strings.Split(line, ",")
		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])
		z, _ := strconv.Atoi(splits[2])

		c := cube{float32(x), float32(y), float32(z)}

		for _, offset := range offsets {
			key := cube{c[0] + offset[0], c[1] + offset[1], c[2] + offset[2]}
			sides[key]++
		}

	}

	sum := 0
	for _, v := range sides {
		if v == 1 {
			sum++
		}
	}

	println(sum)
}
