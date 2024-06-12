package functions

import (
	"fmt"
	"strings"
)

func numirateString(lines *[]string) {

	for index, value := range *lines {
		tmp := append([]string{fmt.Sprintf("%d.", index+1)}, value)
		(*lines)[index] = strings.Join(tmp, " ")
	}

}
