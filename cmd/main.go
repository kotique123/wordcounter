package main

import (
	"fmt"
	"word-counter/pkg/wordcounter"
)

func main() {
	// Output initial phrase
	initial_phrase := "'Joe can't tell between 'large' and large 123 TELL !@#$ programing'"
	fmt.Println(initial_phrase)

	// Invoke formatting sequence
	phrase := initial_phrase
	frequency := wordcounter.FromStringToFrequency(phrase)
	fmt.Print(frequency)
}
