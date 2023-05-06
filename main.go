package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// запрос входной строки у пользователя
	var input string
	fmt.Print("Введите математическое выражение: ")
	fmt.Scanln(&input)

	// разделение строки на операнды и оператор с помощью регулярного выражения.
	re := regexp.MustCompile(`^([IVX]+|\d+)\s*([-+*/])\s*([IVX]+|\d+)$`)
	match := re.FindStringSubmatch(input)
	if len(match) == 0 {
		fmt.Println("Вы ввели неправильное выражение")
		return
	}

	operand1, err := parseOperand(match[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	operand2, err := parseOperand(match[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	operator := match[2]
	result, err := calculate(operand1, operand2, operator)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

// функция проверки операнда на соответствие условиям
func parseOperand(str string) (int, error) {
	n, err := strconv.Atoi(str)
	if err == nil {
		// операнд - арабская цифра
		if n < 1 || n > 10 {
			return 0, fmt.Errorf("операнд должен быть в диапазоне от 1 до 10 включительно")
		}
		return n, nil
	}

	// операнд - римская цифра
	if !isRomanNumeral(str) {
		return 0, fmt.Errorf("недопустимый формат римской цифры")
	}

	n = romanToArabic(str)
	if n < 1 || n > 10 {
		return 0, fmt.Errorf("операнд должен быть в диапазоне от 1 до 10 включительно")
	}

	return n, nil
}

// функция подсчета математического выражения
func calculate(operand1, operand2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return operand1 / operand2, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор")
	}
}

func isRomanNumeral(str string) bool {
	re := regexp.MustCompile(`^[IVX]+$`)
	return re.MatchString(str)
}

func romanToArabic(str string) int {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	var result int
	var prevValue int

	for _, ch := range str {
		value := romanValues[ch]
		if value > prevValue {
			result += value - 2*prevValue
		} else {
			result += value
		}
		prevValue = value
	}

	return result
}
