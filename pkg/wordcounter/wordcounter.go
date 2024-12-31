package wordcounter

import "strings"

func ToLowerCase(phrase string) string {
	return strings.ToLower(phrase)
}

func RemoveQuotes(phrase string) string {
	runes := []rune(phrase)
	switch runes[0] {
	case 39:
		runes[0] = 32
		runes[len(runes)-1] = 32
	}
	return string(runes)
}

func CleanupString(phrase string) string {
	runes := []rune(phrase)
	for i := 0; i < len(phrase); i++ {
		if runes[i] < 97 || runes[i] > 122 && runes[i] != 39 {
			if runes[i]-48 < 0 || runes[i]-48 > 9 {
				runes[i] = 32
			}
		}
	}
	return string(runes)
}

func RemoveApostrophy(phrase string) string {
	if len(phrase) < 1 {
		return phrase
	} else {
		runes := []rune(phrase)
		for i := 0; i < len(phrase)-1; i++ {
			if runes[i] == 39 && phrase[i-1] == 32 {
				runes[i] = 32
			} else if phrase[i] == 39 && phrase[i+1] == 32 {
				runes[i] = 32
			}
		}
		return string(runes)
	}
}

// make sure that the numbers are kept
