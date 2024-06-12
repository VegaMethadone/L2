package functions

import (
	"mansort/structs"
	"sort"
	"strings"
	"time"
)

func monthSort(lines *[]string, args *structs.Args) {
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
			return isMonth(grid[i][index]) > isMonth(grid[j][index])
		})
	} else {
		sort.Slice(grid, func(i, j int) bool {
			return isMonth(grid[i][index]) < isMonth(grid[j][index])
		})
	}
	for i, value := range grid {
		(*lines)[i] = strings.Join(value, " ")
	}
}

func isMonth(month string) int {
	monthNames := []string{
		time.January.String(), time.February.String(), time.March.String(),
		time.April.String(), time.May.String(), time.June.String(),
		time.July.String(), time.August.String(), time.September.String(),
		time.October.String(), time.November.String(), time.December.String(),
	}

	month = strings.ToLower(month)

	for index, m := range monthNames {
		if strings.ToLower(m) == month {
			return index
		}
	}
	return 13
}
