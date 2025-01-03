// Copied from https://github.com/exercism/go/blob/main/exercises/practice/word-count/word_count_test.go
// and edited by Daniil Murashov DaniilMurashov@icloud.com
package wordcounter

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := FromStringToFrequency(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("WordCount(%q)\n got:%v\nwant:%v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkWordCount(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			FromStringToFrequency(tt.input)
		}
	}
}
