package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	games := strings.Split(strings.TrimSpace(input), "\n")

	score := 0

	for _, game := range games {
		signs := strings.Split(game, " ")

		left := signs[0]
		right := signs[1]

		// A = Rock
		// B = Paper
		// C = Scissors

		// X = Rock // Loose
		// Y = Paper // Draw
		// Z = Scissors // Win

		if right == "X" {
			if left == "A" {
				right = "Z"
			} else if left == "B" {
				right = "X"
			} else if left == "C" {
				right = "Y"
			}
		} else if right == "Y" {
			if left == "A" {
				right = "X"
			} else if left == "B" {
				right = "Y"
			} else if left == "C" {
				right = "Z"
			}
		} else if right == "Z" {
			if left == "A" {
				right = "Y"
			} else if left == "B" {
				right = "Z"
			} else if left == "C" {
				right = "X"
			}
		}

		// Win
		if left == "A" && right == "Y" || left == "B" && right == "Z" || left == "C" && right == "X" {
			score += 6
		}

		// Draw
		if left == "A" && right == "X" || left == "B" && right == "Y" || left == "C" && right == "Z" {
			score += 3
		}

		if right == "X" {
			score += 1
		} else if right == "Y" {
			score += 2
		} else if right == "Z" {
			score += 3
		}
	}

	println(score)
}
