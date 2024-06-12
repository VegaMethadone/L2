package functions

import (
	"mansort/structs"
	"sort"
	"strconv"
	"strings"
)

func numericSort(lines *[]string, args *structs.Args) {

	slitLines := func(arr *[]string) [][]string {
		matrix := make([][]string, len(*arr))
		for key := range matrix {
			matrix[key] = strings.Split((*arr)[key], " ")
		}
		return matrix
	}

	grid := slitLines(lines)

	index := args.K - 1

	if args.R {
		sort.Slice(grid, func(i, j int) bool {
			return stoI(grid[i][index]) > stoI(grid[j][index])
		})
	} else {
		sort.Slice(grid, func(i, j int) bool {
			return stoI(grid[i][index]) < stoI(grid[j][index])
		})
	}

	for i, value := range grid {
		(*lines)[i] = strings.Join(value, " ")
	}
}

func stoI(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
