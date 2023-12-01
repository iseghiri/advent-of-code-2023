package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input1 := Input1()
	fmt.Println(input1)
	input2 := Input2()
	fmt.Println(input2)
}

func Input1() int {
	input := ReadInput("/day1/input1")
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	var result int
	for fileScanner.Scan() {
		calibrationValue := BuildCalibrationValue(GetDigit(fileScanner.Text()))
		result += calibrationValue

	}
	input.Close()
	return result
}

func Input2() int {
	input := ReadInput("/day1/input1")
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	var result int
	for fileScanner.Scan() {
		calibrationValue := BuildCalibrationValue2(GetAnyDigit(fileScanner.Text()))
		result += calibrationValue
		fmt.Println("current result after last sum: ", result)
	}
	input.Close()
	return result
}

func GetAnyDigit(line string) []string {
	stringValidDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var collectedSubstr []string
	// Regex to match keys of the string digits map OR digits
	pattern := fmt.Sprintf("(%s)|\\d", strings.Join(getKeys(stringValidDigits), "|"))
	re := regexp.MustCompile(pattern)

	// Find all matches of substrings or digits in the string
	matches := re.FindAllString(line, -1)

	collectedSubstr = append(collectedSubstr, matches...)
	onlyDigits := make([]string, len(collectedSubstr))
	for i, element := range collectedSubstr {
		if num, exists := stringValidDigits[element]; exists {
			onlyDigits[i] = num
		} else {
			onlyDigits[i] = element
		}
	}
	fmt.Println("line :", line)
	fmt.Println("digits:", onlyDigits)
	return onlyDigits
}

func GetDigit(line string) []rune {
	var digits []rune
	for _, char := range line {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}
	return digits
}

// Helper function to extract keys from a map
func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func BuildCalibrationValue(digits []rune) int {
	twoDigitNumber := (string(digits[0]) + string(digits[len(digits)-1]))
	calibrationValue, err := strconv.Atoi(twoDigitNumber)
	check(err)
	return calibrationValue
}

func BuildCalibrationValue2(digits []string) int {
	twoDigitNumber := (string(digits[0]) + string(digits[len(digits)-1]))
	calibrationValue, err := strconv.Atoi(twoDigitNumber)
	check(err)
	fmt.Println("calibrationValue : ", calibrationValue)
	return calibrationValue
}

func ReadInput(path string) *os.File {
	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + path)
	check(err)
	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
