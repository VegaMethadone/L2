package functions

import (
	"mangrep/structs"
	"strconv"
	"strings"
)

func countMatch(lines *[]string, args *structs.Args) []string {
	res := []string{}
	counter := 0
	if args.Fixed {
		for _, str := range *lines {
			if str == args.Arguments[0] {
				counter++
			}
		}
	} else if args.Invert {
		for _, str := range *lines {
			if str != args.Arguments[0] {
				counter++
			}
		}
	} else {
		for _, str := range *lines {
			if strings.Contains(str, args.Arguments[0]) {
				counter++
			}
		}
	}
	res = append(res, strconv.Itoa(counter))
	return res
}
