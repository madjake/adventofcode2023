package main

func main() {
	// day1()
	// day2()
	day3()
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	end := start + length
	if end > len(asRunes) {
		end = len(asRunes)
	}

	return string(asRunes[start:end])
}
