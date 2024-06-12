package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	if unicode.IsDigit(rune(str[0])) {
		return "", errors.New("incorrect string format")
	}

	var builder strings.Builder
	escaped := false

	for _, char := range str {
		if escaped {
			builder.WriteRune(char)
			escaped = false
		} else {
			if char == '\\' {
				escaped = true
			} else if unicode.IsDigit(char) {
				num, _ := strconv.Atoi(string(char))
				lastChar := builder.String()[builder.Len()-1]
				for i := 0; i < num-1; i++ {
					builder.WriteByte(lastChar)
				}

			} else {
				builder.WriteRune(char)
			}
		}
	}

	return builder.String(), nil
}

func main() {

	{
		str := "a4bc2d5e"
		result, err := unpackString(str)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(result)
	}
	{
		str := `qwe\\5`
		result, err := unpackString(str)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(result)
	}
}
