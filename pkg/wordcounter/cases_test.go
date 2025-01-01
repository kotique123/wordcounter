package wordcounter

var testCases = []struct {
	description string
	input       string
	expected    map[string]int
}{
	{
		description: "count one word",
		input:       "word",
		expected:    map[string]int{"word": 1},
	},
	{
		description: "count one of each word",
		input:       "one of each",
		expected:    map[string]int{"each": 1, "of": 1, "one": 1},
	},
	{
		description: "multiple occurrences of a word",
		input:       "one fish two fish red fish blue fish",
		expected:    map[string]int{"blue": 1, "fish": 4, "one": 1, "red": 1, "two": 1},
	},
	{
		description: "handles cramped lists",
		input:       "one,two,three",
		expected:    map[string]int{"one": 1, "three": 1, "two": 1},
	},
	{
		description: "handles expanded lists",
		input:       "one,\ntwo,\nthree",
		expected:    map[string]int{"one": 1, "three": 1, "two": 1},
	},
	{
		description: "ignore punctuation",
		input:       "car: carpet as java: javascript!!&@$%^&",
		expected:    map[string]int{"as": 1, "car": 1, "carpet": 1, "java": 1, "javascript": 1},
	},
	{
		description: "include numbers",
		input:       "testing, 1, 2 testing",
		expected:    map[string]int{"1": 1, "2": 1, "testing": 2},
	},
	{
		description: "normalize case",
		input:       "go Go GO Stop stop",
		expected:    map[string]int{"go": 3, "stop": 2},
	},
	{
		description: "with apostrophes",
		input:       "'First: don't laugh. Then: don't cry. You're getting it.'",
		expected:    map[string]int{"cry": 1, "don't": 2, "first": 1, "getting": 1, "it": 1, "laugh": 1, "then": 1, "you're": 1},
	},
	{
		description: "with quotations",
		input:       "Joe can't tell between 'large' and large.",
		expected:    map[string]int{"and": 1, "between": 1, "can't": 1, "joe": 1, "large": 2, "tell": 1},
	},
	{
		description: "substrings from the beginning",
		input:       "Joe can't tell between app, apple and a.",
		expected:    map[string]int{"a": 1, "and": 1, "app": 1, "apple": 1, "between": 1, "can't": 1, "joe": 1, "tell": 1},
	},
	{
		description: "multiple spaces not detected as a word",
		input:       " multiple   whitespaces",
		expected:    map[string]int{"multiple": 1, "whitespaces": 1},
	},
	{
		description: "alternating word separators not detected as a word",
		input:       ",\n,one,\n ,two \n 'three'",
		expected:    map[string]int{"one": 1, "three": 1, "two": 1},
	},
}
