package day02

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseColour(gameSet string, colour string, defaultValue int) int {
	regex := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, colour))
	match := regex.FindStringSubmatch(gameSet)
	value := defaultValue
	if match != nil {
		value, _ = strconv.Atoi(match[1])
	}
	return value
}

func Part1() int {
	readFile, _ := os.Open("day02/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var part1 int = 0
	for fileScanner.Scan() {
		possibleGame := true
		line := fileScanner.Text()
		re := regexp.MustCompile(`^(Game (\d+): )(.*)`)
		gameId, _ := strconv.Atoi(re.FindStringSubmatch(line)[2])
		lineStripped := re.ReplaceAllString(line, `$3`)
		for _, gameSet := range strings.Split(lineStripped, ";") {
			blue := parseColour(gameSet, "blue", 0)
			red := parseColour(gameSet, "red", 0)
			green := parseColour(gameSet, "green", 0)
			if blue > 14 || red > 12 || green > 13 {
				possibleGame = false
			}
		}
		if possibleGame {
			part1 = part1 + gameId
		}
	}
	return part1
}

func Part2() int {
	readFile, _ := os.Open("day02/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var part2 int = 0
	for fileScanner.Scan() {
		blue, red, green := 0, 0, 0
		possibleGame := true
		line := fileScanner.Text()
		re := regexp.MustCompile(`^(Game (\d+): )(.*)`)
		lineStripped := re.ReplaceAllString(line, `$3`)
		for _, gameSet := range strings.Split(lineStripped, ";") {
			blueN := parseColour(gameSet, "blue", 1)
			redN := parseColour(gameSet, "red", 1)
			greenN := parseColour(gameSet, "green", 1)
			if blueN > blue {
				blue = blueN
			}
			if redN > red {
				red = redN
			}
			if greenN > green {
				green = greenN
			}
		}
		if possibleGame {
			part2 = part2 + red*blue*green
		}
	}
	return part2
}
