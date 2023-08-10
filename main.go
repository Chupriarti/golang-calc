package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("# Welcome to KATA calculator! What to calc?")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	inputParts := strings.Fields(text)

	operand1, err1 := strconv.Atoi(inputParts[0])
	if err1 != nil {
		panic(err1)
	}
	operand2, err2 := strconv.Atoi(inputParts[2])
	if err2 != nil {
		panic(err2)
	}
	operator := inputParts[1]
	var result int

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = int(operand1 / operand2)
	}

	fmt.Println("Result: ")
	fmt.Println(result)
}

func intToRoman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	return roman.String()
}

func romanToInt(s string) int {

	var RomanNumerals = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := RomanNumerals[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}
