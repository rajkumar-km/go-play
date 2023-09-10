package main

import (
	"fmt"
	"unicode/utf8"
)

// DemoUnicode demonstrates using unicode characters in Go
// In earlier days, ASCII characters are used in computers. But this very much limited to
// set of letters used in the world. So, a unicode (unicode.org) is defined to represent
// all the letters.
//
// UTF-32:
//   - UTF-32 is simplest form to hold unicode characters each in 32 bits.
//   - But it occupies much space, because most of the characters for 1 bytes or even 2 bytes.
//   - UTF-8 - a variable length encoding can be more space efficient
//
// UTF-8:
//   - Invented by Ken Thompson and Rob Pike, two of the creators of Go.
//   - In UTF-8, the higher order bits indicates how many bytes to follow
//     0xxxxxxx - 0 indicates ASCII (0-127) stored in remaining 7 bits
//     110xxxxx 10xxxxxx - 2 bytes
//     1110xxxx 10xxxxxx 10xxxxxx - 3 bytes
//     11110xxx 10xxxxxx 10xxxxxx 10xxxxxx - 4 bytes
//   - It is possible to find the first byte from anywhere
//   - Go source files are always encoded in UTF-8
//   - Go string literals can contain unicode characters or their escapes like
//     \u4e16 - 16 bites
//     \U00004e16 - 32 bits with capital U
//     \x41 - 8 bits can be represented with \x anyway
func DemoUnicode() {

	// The UTF-8 allows most of the string properties without decoding
	s := "Hi, தமிழ்"
	fmt.Printf("%s (len=%d, unicodes=%d)\n", s, len(s), utf8.RuneCountInString(s))

	prefix := "த"
	substr := "மிழ்"
	suffix := "ழ்"
	fmt.Println(HasPrefix(s, prefix), Contains(s, substr), HasSuffix(s, suffix))

	// unicode/utf8 package is helpful to process the unicode characters
	fmt.Println(`unicode/utf8 package methods to process unicodes`)
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("  %d: %c\n", i, r)
		i += size
	}

	// The range loop automatically processes the string in UTF-8
	// Go displays � (\uFFFD) character for any malformed UTF-8 characters
	fmt.Println("The range loop also process strings as UTF-8 runes")
	for i, r := range s {
		fmt.Printf("  %d  %q  %[2]d  %[2]x\n", i, r)
	}

	// UTF-8 in string is good for interchange format
	// But, rune types are convinient for programs since they are fixed in size and
	// can be easily indexed in arrays and slices.
	fmt.Println(`rune arrays/slices are convinient for programs`)
	r := []rune(s)
	fmt.Printf("  string = % x\n", s) // "% x" adds space to each characters
	fmt.Printf("  []rune = %x\n", r)
	// converting rune to string produces concatenated UTF-8 encoded string
	fmt.Printf("  %s\n", string(r))

	// Converting integer to rune produces UTF-8 character, � if invalid.
	fmt.Println(`integers can also be converted to string as unicode characters`)
	fmt.Printf("  %s %s %s\n", string(65), string(0xba4), string(82343246))
}

func HasPrefix(s string, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s string, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s string, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
