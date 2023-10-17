package stringhelpers

import "fmt"

func ExampleShorterLength() {
	fmt.Println(ShorterLength("abc", "de"))
	fmt.Println(ShorterLength("ab", "def"))
	fmt.Println(ShorterLength("abc", ""))
	fmt.Println(ShorterLength("", "def"))
	fmt.Println(ShorterLength("abc", "def"))
	fmt.Println(ShorterLength("abc", "def"))
	fmt.Println(ShorterLength("abc", "abc"))

	// Output:
	// 2
	// 2
	// 0
	// 0
	// 3
	// 3
	// 3
}

func ExampleLongerLength() {
	fmt.Println(LongerLength("abc", "de"))
	fmt.Println(LongerLength("ab", "def"))
	fmt.Println(LongerLength("abc", ""))
	fmt.Println(LongerLength("", "def"))
	fmt.Println(LongerLength("abc", "def"))
	fmt.Println(LongerLength("abc", "def"))
	fmt.Println(LongerLength("abc", "abc"))

	// Output:
	// 3
	// 3
	// 3
	// 3
	// 3
	// 3
	// 3
}

func ExampleDifferenceCount() {
	fmt.Println(DifferenceCount("abc", "azc"))
	fmt.Println(DifferenceCount("abc", "zbc"))
	fmt.Println(DifferenceCount("abc", "abz"))
	fmt.Println(DifferenceCount("abc", "abz"))
	fmt.Println(DifferenceCount("abc", "ab"))
	fmt.Println(DifferenceCount("ab", "abc"))
	fmt.Println(DifferenceCount("abcdefg", "a*cd*fg+"))

	// Output:
	// 1
	// 1
	// 1
	// 1
	// 1
	// 1
	// 3
}

func ExamplePositionIsValid() {
	fmt.Println(PositionIsValid("abc", -1))
	fmt.Println(PositionIsValid("abc", 0))
	fmt.Println(PositionIsValid("abc", 1))
	fmt.Println(PositionIsValid("abc", 2))
	fmt.Println(PositionIsValid("abc", 3))

	// Output:
	// false
	// true
	// true
	// true
	// false
}

func ExampleSwitchPositions() {
	fmt.Println(SwitchPositions("abc", 0, 1))
	fmt.Println(SwitchPositions("abc", 0, 2))
	fmt.Println(SwitchPositions("abc", 1, 2))
	fmt.Println(SwitchPositions("abc", 1, 3))

	// Output:
	// bac
	// cba
	// acb
	// abc
}
