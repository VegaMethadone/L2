package functions

import (
	"errors"
	"strconv"
	"strings"
)

func ParseFields(fieldsStr string) ([]int, error) {
	fields := []int{}

	if strings.Contains(fieldsStr, "-") {
		arr := strings.Split(fieldsStr, "-")
		if len(arr) != 2 {
			return nil, errors.New("invalid -f argument: example 1-3")
		}

		start, err := strconv.Atoi(arr[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(arr[1])
		if err != nil {
			return nil, err
		}

		for i := start; i <= end; i++ {
			fields = append(fields, i)
		}

	} else {
		arr := strings.Split(fieldsStr, ",")
		for _, value := range arr {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, errors.New("invalid -f argument: example 1,3 | 1")
			}
			fields = append(fields, num)
		}
	}

	return fields, nil
}
