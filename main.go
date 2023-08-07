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
