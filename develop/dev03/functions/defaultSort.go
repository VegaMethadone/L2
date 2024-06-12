package functions

import (
	"mansort/structs"
	"sort"
	"strings"
)

func defaultSort(lines *[]string, args *structs.Args) {
	if args.K == 0 {
		if args.R {
			sort.Slice(*lines, func(i, j int) bool { return (*lines)[i] > (*lines)[j] })
		} else {
			sort.Slice(*lines, func(i, j int) bool { return (*lines)[i] < (*lines)[j] })
		}
	} else {
		index := args.K - 1

		slitLines := func(arr *[]string) [][]string {
			matrix := make([][]string, len(*arr))
			for key := range matrix {
				matrix[key] = strings.Split((*arr)[key], " ")
			}
			return matrix
		}
		grid := slitLines(lines)

		if args.R {
			sort.Slice(grid, func(i, j int) bool {
				return grid[i][index] > grid[j][index]
			})
		} else {
			sort.Slice(grid, func(i, j int) bool {
				return grid[i][index] < grid[j][index]
			})
		}
		for i, value := range grid {
			(*lines)[i] = strings.Join(value, " ")
		}
	}
}
