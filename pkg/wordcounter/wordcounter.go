// Package wordcounter is a small library that has tools to work with strings
// and count words returning a map.
package wordcounter

import "strings"

// ToLowerCase returns a string in all in lowercase.
func ToLowerCase(phrase string) string {
	return strings.ToLower(phrase)
}

// RemoveQuotes function replaces quotes from beginning and the end of the
// string with a space if they are present and returns a string.
func RemoveQuotes(phrase string) string {
	runes := []rune(phrase)
	switch runes[0] {
	case 39:
		runes[0] = 32
	}
	switch runes[len(runes)-1] {
	case 39:
		runes[len(runes)-1] = 32
	}
	return string(runes)
}

// CleanupString replaces all the special characters from the string with a space except:
// numbers and apostrophies.
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

// RemoveApostrophy removes apostrophies except those that are part of words
// for example, apostrophies in words like "don't" "I'm" "he's" are kept,
// however, if they are surrounding a word, for example 'large', they are being
// replaced with spaces.
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

// CountWords returns a map of strings with int values that contains a word
// and a number of times it ocurrs in a string, including number combinations.
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

// FromStringToFrequency returns a map of string with int values taking any string
// as input and processing it.
func FromStringToFrequency(phrase string) map[string]int {
	phrase = ToLowerCase(phrase)
	phrase = RemoveQuotes(phrase)
	phrase = CleanupString(phrase)
	phrase = RemoveApostrophy(phrase)
	frequency := CountWords(phrase)
	return frequency
}
