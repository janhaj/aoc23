package day01

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Part1() int64 {
	readFile, _ := os.Open("day01/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var part1 int64 = 0
	for fileScanner.Scan() {
		re := regexp.MustCompile(`\s*(\d).*(\d)\s*`)
		match := re.FindStringSubmatch(fileScanner.Text())
		var a int64 = 0
		var b int64 = 0
		if match != nil {
			a, _ = strconv.ParseInt(match[1], 10, 0)
			b, _ = strconv.ParseInt(match[2], 10, 0)
		} else {
			re = regexp.MustCompile(`.*(\d).*`)
			match := re.FindStringSubmatch(fileScanner.Text())
			a, _ = strconv.ParseInt(match[1], 10, 0)
			b = a
		}
		part1 = part1 + a*10 + b
	}
	return part1
}

func replace(value string, err error) int64 {
	myDict := make(map[string]int)

	// Add key-value pairs to the map
	myDict["one"] = 1
	myDict["two"] = 2
	myDict["three"] = 3
	myDict["four"] = 4
	myDict["five"] = 5
	myDict["six"] = 6
	myDict["seven"] = 7
	myDict["eight"] = 8
	myDict["nine"] = 9
	return int64(myDict[value])
}

func Part2() int64 {
	readFile, _ := os.Open("day01/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var part2 int64 = 0
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		re := regexp.MustCompile(`\s*(one|two|three|four|five|six|seven|eight|nine|\d).*(one|two|three|four|five|six|seven|eight|nine|\d)\s*`)
		match := re.FindStringSubmatch(fileScanner.Text())
		var a int64 = 0
		var b int64 = 0
		var err error
		if match != nil {
			a, err = strconv.ParseInt(match[1], 10, 0)
			if err != nil {
				a = replace(match[1], err)
			}
			b, err = strconv.ParseInt(match[2], 10, 0)
			if err != nil {
				b = replace(match[2], err)
			}
		} else {
			re = regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|\d).*`)
			match := re.FindStringSubmatch(fileScanner.Text())
			a, err = strconv.ParseInt(match[1], 10, 0)
			if err != nil {
				a = replace(match[1], err)
			}
			b = a
		}
		part2 = part2 + a*10 + b
	}
	return part2
}
