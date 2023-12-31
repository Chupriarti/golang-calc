package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ката Калькулятор. Поддерживаются операции с арабскими и римскими числами: +, -, *, / .")
	fmt.Println("Введитие выражение в формате: 4 + 7 или XIV * X")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	inputParts := strings.Fields(text)
	if len(inputParts) != 3 {
		panic("ОШИБКА: Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	operand1, err1 := strconv.Atoi(inputParts[0])
	operand2, err2 := strconv.Atoi(inputParts[2])
	isRomanCalculation := err1 != nil && err2 != nil
	operator := inputParts[1]
	var result int

	if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) {
		panic("ОШИБКА: Используются одновременно разные системы счисления")
	}

	if isRomanCalculation {
		operand1 = romanToInt(inputParts[0])
		operand2 = romanToInt(inputParts[2])
	}

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		if isRomanCalculation && operand2 > operand1 {
			panic("ОШИБКА: В римской системе нет отрицательных чисел")
		}
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = int(operand1 / operand2)
	}

	if isRomanCalculation {
		fmt.Println("Результат:", intToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}

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
