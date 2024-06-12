package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortWord(str string) string {
	newRune := []rune(str)
	sort.Slice(newRune, func(i, j int) bool { return newRune[i] < newRune[j] })
	return string(newRune)
}

func FindAnagrams(arr *[]string) *map[string][]string {
	for i, word := range *arr {
		(*arr)[i] = strings.ToLower(word)
	}
	seen := make(map[string]bool)
	m := make(map[string][]string)

	for _, word := range *arr {
		if len(word) == 1 {
			continue
		}
		if seen[word] {
			continue
		} else {
			seen[word] = true
			sortedWord := sortWord(word)
			m[sortedWord] = append(m[sortedWord], word)
		}
	}
	result := make(map[string][]string)

	for _, slice := range m {
		if len(slice) < 2 {
			continue
		} else {
			firstWord := slice[0]
			sort.Strings(slice)
			tmp := make([]string, len(slice))
			copy(tmp, slice)
			result[firstWord] = tmp
		}
	}

	return &result
}

func main() {
	arr := []string{"пятак", "пятка", "тяпка", "тяпка", "листок", "листок", "слиток", "столик"}
	m := FindAnagrams(&arr)

	for key, value := range *m {
		fmt.Println(key, "\t", value)
	}

}
