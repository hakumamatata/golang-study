package types

import "strings"

func JoinWithCommas(phrases []string) string {
	var result string
	if len(phrases) == 0 {
		result = ""
	} else if len(phrases) == 1 {
		result = phrases[0]
	} else if len(phrases) == 2 {
		result = phrases[0] + " and " + phrases[1]
	} else {
		result = strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
	}
	return result
}
