package wordcounter

import "strings"

func ToLowerCase(phrase string) string {
	return strings.ToLower(phrase)
}

func RemoveQuotes(phrase string) []rune {
	runes := []rune(phrase)
	switch runes[0] {
	case 39:
		runes[0] = 32
		runes[len(runes)-1] = 32
	}
	return runes
}

func CleanupString(phrase []rune) string {
	clean_string := strings.Trim(string(phrase), "!&@$%^#()\"+-`~.")
	return clean_string
}

func RemoveApostrophy(phrase []rune) []rune {
	if len(phrase) < 1 {
		return phrase
	} else {
		for i := 0; i < len(phrase)-1; i++ {
			if phrase[i] == 39 && phrase[i-1] == 32 {
				phrase[i] = 32
			} else if phrase[i] == 39 && phrase[i+1] == 32 {
				phrase[i] = 32
			}
		}
	}
	return phrase
}

// make sure that the numbers are kept
