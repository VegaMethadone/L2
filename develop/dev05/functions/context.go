package functions

import "mangrep/structs"

func findContext(lines *[]string, args *structs.Args) []string {
	args.Before = args.Context
	args.After = args.Context

	res1 := findBefore(lines, args)
	res2 := findAfter(lines, args)

	return append(res1, res2...)
}
