package functions

import (
	"mangrep/structs"
	"strings"
)

func findBefore(lines *[]string, args *structs.Args) []string {
	res := []string{}
	counter := args.Before

	if args.Fixed {
		for index, str := range *lines {
			if str == args.Arguments[0] {
				i := index - 1
				for i >= 0 && counter > 0 {
					res = append(res, (*lines)[i])
					counter--
					i--
				}
				break
			}
		}

	} else if args.Invert {
		for index, str := range *lines {
			if str != args.Arguments[0] {
				i := index - 1
				for i >= 0 && counter > 0 {
					res = append(res, (*lines)[i])
					counter--
					i--
				}
				break
			}
		}

	} else {
		for index, str := range *lines {
			if strings.Contains(str, args.Arguments[0]) {
				i := index - 1
				for i >= 0 && counter > 0 {
					res = append(res, (*lines)[i])
					counter--
					i--
				}
				break
			}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
