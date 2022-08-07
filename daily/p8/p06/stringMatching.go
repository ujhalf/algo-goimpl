package p06

import "strings"

func stringMatching(words []string) []string {
	var ret []string
	for i, s := range words {
		for j, t := range words {
			if i != j && strings.Contains(t, s) {
				ret = append(ret, s)
				break
			}
		}
	}
	return ret
}
