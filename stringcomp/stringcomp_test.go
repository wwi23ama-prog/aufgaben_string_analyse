package stringcomp

import "fmt"

func ExampleAtMostNDifferences() {
	fmt.Println(AtMostNDifferences("abc", "azc", 0))
	fmt.Println(AtMostNDifferences("abc", "azc", 1))
	fmt.Println(AtMostNDifferences("abc", "azc", 2))

	fmt.Println(AtMostNDifferences("abc", "abc", 0))
	fmt.Println(AtMostNDifferences("abc", "abc", 1))
	fmt.Println(AtMostNDifferences("abc", "abc", 2))

	fmt.Println(AtMostNDifferences("abc", "def", 2))
	fmt.Println(AtMostNDifferences("abc", "def", 3))
	fmt.Println(AtMostNDifferences("abc", "def", 4))

	fmt.Println(AtMostNDifferences("abc", "abcde", 1))
	fmt.Println(AtMostNDifferences("abc", "abcde", 2))
	fmt.Println(AtMostNDifferences("abc", "abcde", 2))

	fmt.Println(AtMostNDifferences("abc", "", 2))
	fmt.Println(AtMostNDifferences("abc", "", 3))

	// Output:
	// false
	// true
	// true
	// true
	// true
	// true
	// false
	// true
	// true
	// false
	// true
	// true
	// false
	// true
}

func ExampleStartsWith() {
	fmt.Println(StartsWith("abc", "a"))
	fmt.Println(StartsWith("abc", "b"))
	fmt.Println(StartsWith("abc", "ab"))
	fmt.Println(StartsWith("abc", "bc"))
	fmt.Println(StartsWith("abc", "abcd"))
	fmt.Println(StartsWith("abc", ""))

	// Output:
	// true
	// false
	// true
	// false
	// false
	// true
}

func ExampleEndsWith() {
	fmt.Println(EndsWith("abc", "c"))
	fmt.Println(EndsWith("abc", "b"))
	fmt.Println(EndsWith("abc", "bc"))
	fmt.Println(EndsWith("abc", "ab"))
	fmt.Println(EndsWith("abc", "abcd"))
	fmt.Println(EndsWith("abc", ""))

	// Output:
	// true
	// false
	// true
	// false
	// false
	// true
}

func ExampleContains() {
	fmt.Println(Contains("abc", ""))
	fmt.Println(Contains("abc", "a"))
	fmt.Println(Contains("abc", "b"))
	fmt.Println(Contains("abc", "c"))
	fmt.Println(Contains("abc", "ab"))
	fmt.Println(Contains("abc", "bc"))
	fmt.Println(Contains("abc", "ac")) // false
	fmt.Println(Contains("abc", "abc"))
	fmt.Println(Contains("abc", "abcd")) // false

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// true
	// false
}

func ExampleEqualCaseInsensitive() {
	fmt.Println(EqualCaseInsensitive("abc", "abc"))
	fmt.Println(EqualCaseInsensitive("abc", "ABC"))
	fmt.Println(EqualCaseInsensitive("abc", "AbC"))
	fmt.Println(EqualCaseInsensitive("abc", "AbCd")) // false

	// Output:
	// true
	// true
	// true
	// false
}

func ExampleEqualExceptTransposedPositions() {
	fmt.Println(EqualExceptTransposedPositions("abc", "abc")) // false
	fmt.Println(EqualExceptTransposedPositions("abc", "acb"))
	fmt.Println(EqualExceptTransposedPositions("abc", "bac"))
	fmt.Println(EqualExceptTransposedPositions("abc", "cba"))         // false
	fmt.Println(EqualExceptTransposedPositions("abc", "abcd"))        // false
	fmt.Println(EqualExceptTransposedPositions("abc", ""))            // false
	fmt.Println(EqualExceptTransposedPositions("abcdefg", "acbdfeg")) // false
	fmt.Println(EqualExceptTransposedPositions("abcdefg", "abcdfeg"))

	// Output:
	// false
	// true
	// true
	// false
	// false
	// false
	// false
	// true
}

func ExampleContainsScatteredSubstring() {
	// Basisfälle
	fmt.Println(ContainsScatteredSubstring("abc", ""))
	fmt.Println(ContainsScatteredSubstring("abc", "a"))
	fmt.Println(ContainsScatteredSubstring("abc", "b"))
	fmt.Println(ContainsScatteredSubstring("abc", "c"))
	fmt.Println(ContainsScatteredSubstring("abc", "ab"))
	fmt.Println(ContainsScatteredSubstring("abc", "bc"))
	fmt.Println(ContainsScatteredSubstring("abc", "ac"))
	fmt.Println(ContainsScatteredSubstring("abc", "abc"))
	fmt.Println(ContainsScatteredSubstring("abc", "abcd")) // false

	// Buchstabendreher (kein verstreuter Teilstring)
	fmt.Println(ContainsScatteredSubstring("abc", "acb"))
	fmt.Println(ContainsScatteredSubstring("abc", "bac"))
	fmt.Println(ContainsScatteredSubstring("abc", "acbd"))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}
