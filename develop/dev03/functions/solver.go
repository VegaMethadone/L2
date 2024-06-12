package functions

import (
	"mansort/structs"
)

func Solver(lines []string, args *structs.Args) []string {
	if args.U {
		unique(&lines)
	}

	if args.M {
		monthSort(&lines, args)
	} else if args.N {
		numericSort(&lines, args)
	} else {
		defaultSort(&lines, args)
	}
	return lines
}
