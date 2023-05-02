package citydigit

import "gopackages/wordz"

func Digit() string {
	wordz.Words = []string{"one", "two", "three", "four", "five"}
	digit := wordz.Random()
	return digit
}
