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
		if (runes[i] < 97 || runes[i] > 122) && runes[i] != 39 {
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
		for i := 0; i < len(runes)-1; i++ {
			if runes[i] == 39 && runes[i-1] == 32 {
				runes[i] = 32
			} else if runes[i] == 39 && runes[i+1] == 32 {
				runes[i] = 32
			}
		}
		if runes[len(runes)-1] == 39 {
			runes[len(runes)-1] = 32
		}
		return string(runes)
	}
}

func CountWords(phrase string) map[string]int {
	words := strings.Fields(phrase)
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word] = 0
	}
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func FromStringToFrequency(phrase string) map[string]int {
	phrase = ToLowerCase(phrase)
	phrase = RemoveQuotes(phrase)
	phrase = CleanupString(phrase)
	phrase = RemoveApostrophy(phrase)
	frequency := CountWords(phrase)
	return frequency
}
