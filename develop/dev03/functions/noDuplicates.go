package functions

func unique(lines *[]string) {
	seen := make(map[string]bool)
	uniqueLines := []string{}

	for _, line := range *lines {
		if seen[line] {
			continue
		} else {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	*lines = uniqueLines
}
