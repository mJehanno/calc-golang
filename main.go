package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("first number ?")
	input1, _ := reader.ReadString('\n')
	input1 = strings.Replace(input1, "\n", "", -1)
	number1, _ := strconv.ParseFloat(input1, 32)
	fmt.Println("operation ?")
	operation, _ := reader.ReadString('\n')
	operation = strings.Replace(operation, "\n", "", -1)
	fmt.Println("second number ?")
	input2, _ := reader.ReadString('\n')
	input2 = strings.Replace(input2, "\n", "", -1)
	number2, _ := strconv.ParseFloat(input2, 32)

	var result float32
	var err error
	switch operation {
	case "sum", "addition", "+":
		result = sum(float32(number1), float32(number2))
	case "sub", "minus", "substraction", "-":
		result = substraction(float32(number1), float32(number2))
	case "multiplication", "*":
		result = multiplication(float32(number1), float32(number2))
	case "division", "/":
		result, err = division(float32(number1), float32(number2))
	}

	if err == nil {

		fmt.Println("result is : " + fmt.Sprintf("%f", result))
	} else {
		fmt.Println(fmt.Errorf("Error : " + err.Error()))
	}

}

func sum(a float32, b float32) float32 {
	return a + b
}

func substraction(a float32, b float32) float32 {
	return a - b
}

func multiplication(a float32, b float32) float32 {
	return a * b
}

func division(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return a / b, nil
}
