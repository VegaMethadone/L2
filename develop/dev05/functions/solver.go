package functions

import (
	"mangrep/structs"
	"strings"
)

func TheGod(lines *[]string, args *structs.Args) []string {
	result := []string{}

	if args.Linenum {
		numirateString(lines)
	}
	if args.Ignore {
		ignoreCase(lines)
		args.Arguments[0] = strings.ToLower(args.Arguments[0])
	}

	if args.Before > 0 {
		result = append(result, findBefore(lines, args)...)
	}
	if args.After > 0 {
		result = append(result, findAfter(lines, args)...)
	}
	if args.Context > 0 {
		result = append(result, findContext(lines, args)...)
	}

	if args.Count {
		result = append(result, countMatch(lines, args)...)
	}

	return result
}
