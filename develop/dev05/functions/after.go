package functions

import (
	"mangrep/structs"
	"strings"
)

func findAfter(lines *[]string, args *structs.Args) []string {
	res := []string{}
	counter := args.After

	if args.Fixed {
		for index, str := range *lines {
			if str == args.Arguments[0] {
				i := index + 1
				for i < len(*lines) && counter > 0 {
					res = append(res, (*lines)[i])
					counter--
					i++
				}
				break
			}
		}

	} else if args.Invert {
		for index, str := range *lines {
			if str != args.Arguments[0] {
				i := index + 1
				for i < len(*lines) && counter > 0 {
					res = append(res, (*lines)[i])
					counter--
					i++
				}
				break
			}
		}
	} else {
		for index, str := range *lines {
			if strings.Contains(str, args.Arguments[0]) {
				i := index + 1
				for i < len(*lines) && counter > 0 {
					res = append(res, (*lines)[i])
					counter--
					i++
				}
				break
			}
		}
	}
	return res
}
