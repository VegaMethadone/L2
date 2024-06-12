package functions

import (
	"strings"
)

func ignoreCase(lines *[]string) {

	partLines := func(arr []string) [][]string {
		matrix := make([][]string, len(arr))
		for index := range arr {
			matrix[index] = strings.Split(arr[index], " ")
		}
		return matrix
	}

	grid := partLines(*lines)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = strings.ToLower(grid[i][j])
		}
	}

	for key, values := range grid {
		(*lines)[key] = strings.Join(values, " ")
	}

}
