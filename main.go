package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arabicToRoman(num int) string {
	if num < 1 {
		panic("There are no negative numbers in the Roman system")
	}

	romanNumerals := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}

	keys := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string

	for _, key := range keys {
		for num >= key {
			result = result + romanNumerals[key]
			num -= key
		}
	}

	return result
}
func romanToArabic(roman string) int {

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	prev := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value, err := romanMap[roman[i]]
		if !err {
			panic("Incorrect input")
		}
		if value < prev {
			result -= value
		} else {
			result += value
		}
		prev = value
	}

	return result
}

func validateNumbers(num1, num2 int) bool {
	if (num1 >= 1 && num1 <= 10) && (num2 >= 1 && num2 <= 10) {
		return true
	}
	panic("The calculator must accept numbers from 1 to 10 inclusive, no more.")
}

func operation(num1Int, num2Int int, operator string) int {
	var result int
	switch operator {
	case "+":
		result = num1Int + num2Int
	case "-":
		result = num1Int - num2Int
	case "*":
		result = num1Int * num2Int
	case "/":
		result = num1Int / num2Int
	default:
		panic("Invalid Operator")
	}
	return result
}

func calculate(input string) (int, bool) {
	items := strings.Fields(input)
	if len(items) != 3 {
		panic("Invalid input. Please provide an expression with three elements.")
	}
	num1 := items[0]
	operator := items[1]
	num2 := items[2]

	if num1Int, err := strconv.Atoi(num1); err == nil {
		num2Int, err := strconv.Atoi(num2)
		if err != nil {
			panic("Incorrect input")
		}
		if validateNumbers(num1Int, num2Int) {
			return operation(num1Int, num2Int, operator), false
		}
	} else {
		num1Int := romanToArabic(num1)
		num2Int := romanToArabic(num2)

		if validateNumbers(num1Int, num2Int) {
			return operation(num1Int, num2Int, operator), true
		}
	}
	return 0, false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	result, flag := calculate(input)
	if !flag {
		fmt.Println(result)
	} else {
		fmt.Println(arabicToRoman(result))
	}

}
