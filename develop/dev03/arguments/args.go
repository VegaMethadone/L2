package arguments

import (
	"errors"
	"flag"
	"mansort/structs"
)

func GetFlags() (*structs.Args, error) {
	k := flag.Int("k", 0, "define on witch column apply sort")
	n := flag.Bool("n", false, "sort by numeric value")
	r := flag.Bool("r", false, "sort in reverse order")
	u := flag.Bool("u", false, "do not output duplicate lines")
	M := flag.Bool("M", false, "sort by month")

	flag.Parse()

	args := &structs.Args{
		K: *k,
		N: *n,
		R: *r,
		U: *u,
		M: *M,
	}

	if args.K < 0 {
		return nil, errors.New("invalid flag -k: less than 0")
	}

	args.Files = flag.Args()
	return args, nil
}
