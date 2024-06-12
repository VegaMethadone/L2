package arguments

import (
	"errors"
	"flag"
	"mangrep/structs"
)

func GetArgs() (*structs.Args, error) {
	A := flag.Int("A", 0, "print +N lines after a match")
	B := flag.Int("B", 0, "print +N lines before a match")
	C := flag.Int("C", 0, "(A+B) print ±N lines around the match")
	c := flag.Bool("c", false, "count lines")
	i := flag.Bool("i", false, "ignore case")
	v := flag.Bool("v", false, "exclude")
	F := flag.Bool("F", false, "exact match with the string")
	n := flag.Bool("n", false, "print the line number")

	flag.Parse()

	args := &structs.Args{
		After:   *A,
		Before:  *B,
		Context: *C,
		Count:   *c,
		Ignore:  *i,
		Invert:  *v,
		Fixed:   *F,
		Linenum: *n,
	}

	args.Arguments = flag.Args()
	if args.After < 0 {
		return nil, errors.New("invalid argument for -A")
	}
	if args.Before < 0 {
		return nil, errors.New("invalid argument for -B")
	}
	if args.Context < 0 {
		return nil, errors.New("invalid argument for -C")
	}
	if args.After > 0 && args.Count || args.Before > 0 && args.Count || args.Context > 0 && args.Count {
		return nil, errors.New("invalid argument:  no mathc with  -A || -B || -C with -c")
	}
	if args.Invert && args.Fixed {
		return nil, errors.New("invalid argument: no match with -v && -F")
	}
	if args.Count && args.Linenum {
		return nil, errors.New("invalid argument: no match with -c && -n")
	}
	if args.Context > 0 && args.After > 0 || args.Context > 0 && args.Before > 0 {
		return nil, errors.New("invalid argument: no match with -C && -A || -B")
	}

	return args, nil
}
