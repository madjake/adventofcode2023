package main

import (
	"fmt"
	"strings"
)

const (
	day6TestInput = `Time:      7  15   30
Distance:  9  40  200`

	day6FullInput = `Time:        58     81     96     76
Distance:   434   1041   2219   1218`

	day6Part2TestInput = `Time:      71530
Distance:  940200`

	day6Part2FullInput = `Time:      58819676
Distance:  434104122191218`
)

/*
*

	Toy boat:
		Starting speed: 0 mm/s
		Speed increase +1 mm/s button is held but stationary until you let go

		button hold ms * remaining ms
*/
type day6RaceData struct {
	Times     []numberMatch
	Distances []numberMatch
}

func day6() {
	result := 0
	result = day6Part1(day6TestInput)

	fmt.Printf("PART 1: Total ways to win (test input): %d\n", result)
	expected := 288
	if result != expected {
		panic(fmt.Sprintf("Part 1 test result invalid. Expected %d but got %d", expected, result))
	}

	result = day6Part1(day6FullInput)

	fmt.Printf("PART 1: Total ways to win (FULL input): %d\n", result)
	expected = 1159152

	if result != expected {
		panic(fmt.Sprintf("Part 1 FULL result invalid. Expected %d but got %d", expected, result))
	}

	result = day6Part2(day6Part2TestInput)
	fmt.Printf("PART 2: Total ways to win (Part 2 Test input): %d\n", result)
	expected = 71503
	if result != expected {
		panic(fmt.Sprintf("Part 2 test result invalid. Expected %d but got %d", expected, result))
	}

	result = day6Part2(day6Part2FullInput)
	fmt.Printf("PART 2: Total ways to win (Part 2 Test input): %d\n", result)
	expected = 41513103
	if result != expected {
		panic(fmt.Sprintf("Part 2 test result invalid. Expected %d but got %d", expected, result))
	}

}

func day6Part1(input string) int {
	raceData := parseDay6Input(input)

	waysToWin := []int{}
	totalWaysToWin := 1

	for i := range raceData.Times {
		waysToWin = append(waysToWin, day6waysToWinRace(raceData.Times[i].num, raceData.Distances[i].num))
	}

	// fmt.Print(raceData, waysToWin)

	for _, c := range waysToWin {
		totalWaysToWin *= c
	}
	return totalWaysToWin
}

func day6Part2(input string) int {
	raceData := parseDay6Input(input)

	return day6waysToWinRace(raceData.Times[0].num, raceData.Distances[0].num)
}

func day6waysToWinRace(time int, distance int) int {
	/*
	   	day6TestInput = `Time:      7  15   30
	   Distance:  9  40  200`
	*/

	// beat 9 mm in 7ms
	// trade 1ms for 1mm of mm/s

	wins := 0

	for i := 0; i < time; i++ {
		speed := i
		d := speed * (time - i)
		if d > distance {
			wins++
			// fmt.Printf("Hold button for %d ms to go %d mm\n", i, d)
		}
	}
	return wins
}

func parseDay6Input(input string) day6RaceData {
	lines := strings.Split(input, "\n")

	times := parseNumbersFromLine(lines[0])
	distances := parseNumbersFromLine(lines[1])

	return day6RaceData{
		Times:     times,
		Distances: distances,
	}
}
