package structs

type Args struct {
	After     int
	Before    int
	Context   int
	Count     bool
	Ignore    bool
	Invert    bool
	Fixed     bool
	Linenum   bool
	Arguments []string
}
