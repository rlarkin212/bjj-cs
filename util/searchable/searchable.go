package searchable

import "strings"

const MIN_SEARCH_LEN = 3

func GenerateSearchable(value string) string {
	out := []string{}

	for _, word := range strings.Split(value, " ") {
		chars := []rune(word)

		for i := range chars {
			left := word[:i]
			if len(left) >= MIN_SEARCH_LEN {
				out = append(out, left)
			}

			right := word[i:]
			if len(right) >= MIN_SEARCH_LEN {
				out = append(out, right)
			}
		}
	}

	res := strings.Join(out, (" "))

	return res
}
