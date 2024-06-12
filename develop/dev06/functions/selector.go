package functions

func SelectColumn(columns []string, fields []int) []string {
	selectedColumns := []string{}

	for _, field := range fields {
		if field > 0 && field <= len(columns) {
			selectedColumns = append(selectedColumns, columns[field-1])
		}
	}

	return selectedColumns
}
