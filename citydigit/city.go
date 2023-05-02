package citydigit

import "gopackages/wordz"

func City() string {
	wordz.Words = []string{"msc", "pnz", "omsk", "prm", "tmbv"}
	city := wordz.Random()
	return city
}
