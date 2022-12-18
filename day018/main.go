package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type cube [3]float64

var offsets = []cube{{0.5, 0, 0}, {-0.5, 0, 0}, {0, 0.5, 0}, {0, -0.5, 0}, {0, 0, 0.5}, {0, 0, -0.5}}

func main() {
	//part1()
	part2()
}

func part1() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sides := make(map[cube]int)

	for _, line := range lines {
		splits := strings.Split(line, ",")
		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])
		z, _ := strconv.Atoi(splits[2])

		c := cube{float64(x), float64(y), float64(z)}

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

func part2() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sides := make(map[cube]int)

	minCube := cube{float64(1000), float64(1000), float64(1000)}
	maxCube := cube{float64(-1000), float64(-1000), float64(-1000)}

	cubes := make(map[cube]bool)

	for _, line := range lines {
		splits := strings.Split(line, ",")
		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])
		z, _ := strconv.Atoi(splits[2])

		c := cube{float64(x), float64(y), float64(z)}

		cubes[c] = true

		for _, offset := range offsets {
			key := cube{c[0] + offset[0], c[1] + offset[1], c[2] + offset[2]}
			sides[key]++
		}

		// Calculate grid size
		minCube[0] = min(minCube[0], c[0])
		minCube[1] = min(minCube[1], c[1])
		minCube[2] = min(minCube[2], c[2])

		maxCube[0] = max(maxCube[0], c[0])
		maxCube[1] = max(maxCube[1], c[1])
		maxCube[2] = max(maxCube[2], c[2])
	}

	minCube[0]--
	minCube[1]--
	minCube[2]--

	maxCube[0]++
	maxCube[1]++
	maxCube[2]++

	airCubes := make(map[cube]bool)
	airCubes[minCube] = true

	queue := []cube{maxCube}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, offset := range offsets {
			pos := cube{current[0] + offset[0] + offset[0],
				current[1] + offset[1] + offset[1],
				current[2] + offset[2] + offset[2]}

			if minCube[0] <= pos[0] && pos[0] <= maxCube[0] &&
				minCube[1] <= pos[1] && pos[1] <= maxCube[1] &&
				minCube[2] <= pos[2] && pos[2] <= maxCube[2] {

				if cubes[pos] || airCubes[pos] {
					continue
				}

				airCubes[pos] = true
				queue = append(queue, pos)
			} else {
				continue
			}
		}

	}

	intersections := make(map[cube]bool)
	for c := range airCubes {
		for _, offset := range offsets {
			intersections[cube{c[0] + offset[0], c[1] + offset[1], c[2] + offset[2]}] = true
		}
	}

	sum := 0
	for s := range sides {
		if intersections[s] {
			sum++
		}
	}

	println(sum)
}

func min(left, right float64) float64 {
	if left < right {
		return left
	}
	return right
}

func max(left, right float64) float64 {
	if left > right {
		return left
	}
	return right
}
