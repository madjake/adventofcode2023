package main

import (
	"regexp"
	"strconv"
)

type numberMatch struct {
	num int
	x   int
	y   int
}

type runeMatch struct {
	c rune
	x int
	y int
}

func parseNumbersFromLine(line string) []numberMatch {
	re := regexp.MustCompile(`\d+`)

	numbers := []numberMatch{}

	matches := re.FindAllStringIndex(line, -1)

	for _, match := range matches {
		num, _ := strconv.Atoi(line[match[0]:match[1]])
		numbers = append(numbers, numberMatch{
			num: num,
			x:   match[0],
			y:   0,
		})
	}

	return numbers
}

func matchNumbers(lines []string) []numberMatch {
	re := regexp.MustCompile(`\d+`)

	numbers := []numberMatch{}

	for y, line := range lines {
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			num, _ := strconv.Atoi(line[match[0]:match[1]])
			numbers = append(numbers, numberMatch{
				num: num,
				x:   match[0],
				y:   y,
			})
		}
	}

	return numbers
}

func matchRunes(r rune, lines []string) []runeMatch {
	runes := []runeMatch{}

	for y, line := range lines {
		for x, c := range line {
			if c == r {
				runes = append(runes, runeMatch{
					c: c,
					y: y,
					x: x,
				})
			}
		}
	}

	return runes
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
