package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/iseghiri/advent-of-code-2023/internal/utils"
)

func main() {
	//FirstStar()
	SecondStar()
}

func SecondStar() {
	input := utils.ReadInput("/day2/input")
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	result := 0
	colors := [3]string{"blue", "red", "green"}
	var patternStrings []string
	for _, color := range colors {
		patternStrings = append(patternStrings, fmt.Sprintf("\\d+ %s", color))
	}
	pattern := strings.Join(patternStrings, "|")
	re := regexp.MustCompile(pattern)
	for fileScanner.Scan() {
		colorsCounter := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}
		matches := re.FindAllString(fileScanner.Text(), -1)
		for _, match := range matches {
			parts := strings.Split(match, " ")
			number, err := strconv.Atoi(parts[0])
			utils.Check(err)
			color := parts[1]
			if number > colorsCounter[color] {
				colorsCounter[color] = number
			}

		}
		product := 1
		for _, maxNumber := range colorsCounter {
			product = product * maxNumber
		}
		result += product

	}
	input.Close()
	println("result : ", result)

}

func FirstStar() {
	input := utils.ReadInput("/day2/input")
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	result := 0
	colors := [3]string{"blue", "red", "green"}
	var patternStrings []string
	for _, color := range colors {
		patternStrings = append(patternStrings, fmt.Sprintf("\\d+ %s", color))
	}
	pattern := strings.Join(patternStrings, "|")
	fmt.Println("pattern : ", pattern)
	re := regexp.MustCompile(pattern)
	for fileScanner.Scan() {
		isValid := true
		colorsCounter := map[string]int{
			"blue":  14,
			"red":   12,
			"green": 13,
		}
		line := fileScanner.Text()
		gameId := GetGameId(fileScanner.Text())
		fmt.Println("GameId : ", gameId)
		matches := re.FindAllString(fileScanner.Text(), -1)
		fmt.Println("matches", matches)
		for _, match := range matches {
			parts := strings.Split(match, " ")
			number, err := strconv.Atoi(parts[0])
			utils.Check(err)
			color := parts[1]
			if number > colorsCounter[color] {
				isValid = false
			}
			if !isValid {
				fmt.Println("colorsCounter : ", colorsCounter)
				fmt.Println("(not valid) Line: ", line)
				break
			}
		}
		if isValid {
			fmt.Println("colorsCounter : ", colorsCounter)
			fmt.Println("(valid) Line :", line)
			result += gameId
		}
	}
	input.Close()
	println("result : ", result)
}

func GetGameId(line string) int {
	gameAndId := strings.Split(strings.Split(line, ":")[0], " ")
	gameID, err := strconv.Atoi(gameAndId[1])
	utils.Check(err)
	return gameID
}

// Helper function to extract keys from a map
func getKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
