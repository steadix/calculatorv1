package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Честно скопированная функция конвертации из римских чисел в арабские
func romanToInt(s string) int {

	characterMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}
	length := len(s)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return characterMap[s[0]]
	}
	sum := characterMap[s[length-1]]
	for i := length - 2; i >= 0; i-- {
		if characterMap[s[i]] < characterMap[s[i+1]] {
			sum -= characterMap[s[i]]
		} else {
			sum += characterMap[s[i]]
		}
	}
	return sum
}

// конвертация из арабских в римские числа
func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func main() {

	var a int           // тут будет храниться первое число
	var b int           // тут будет храниться второе число
	var operator string // тут будет храниться знак
	var input string    // сюда вводится выражение
	var con []string    // сюда нарезаем первое, второе число и знак
	var result int      // тут будет храниться результат вычисления

	// считываем текст, убираем лишнюю табуляцию, пробелы, нарезаем в слайсы
	fmt.Println("Введите выражение:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	con = strings.Fields(input)

	// Конвертируем текст в числа, если числа римские, то временно конвертируем их в арабские.
	// Добавим счетчик, который покажет, сколько из двух введенных чисел являются римскими.
	counter := 0
	a, err1 := strconv.Atoi(con[0])
	if err1 != nil {
		a = romanToInt(con[0])
		counter += 1
	}
	b, err2 := strconv.Atoi(con[2])
	if err2 != nil {
		b = romanToInt(con[2])
		counter += 1
	}
	operator = con[1]

	// Проверяем диапазон, правильность введенного знака и совпадает ли тип введенных чисел.
	// Если счетчик равен 1, значит одно число римское, а другое арабское, тогда будет ошибка.
	if counter != 1 {
		if a > 0 && a <= 10 && b > 0 && b <= 10 {
			if operator == "+" {
				result = a + b
			} else if operator == "-" {
				result = a - b
			} else if operator == "*" {
				result = a * b
			} else if operator == "/" {
				result = a / b
			} else {
				fmt.Println("Ошибка при вводе знака!")
			}
		} else {
			fmt.Println("Работаю только с арабскими и римскими числами от 1 до 10!")
		}
	} else {
		fmt.Println("Введенные числа должны быть одного типа!")
	}
	//Проверяем римское выражение на отрицательность и конвертируем результат обратно в римское число
	if counter == 2 && a > 0 && a <= 10 && b > 0 && b <= 10 {
		if result < 1 {
			fmt.Println("Римские числа не могут быть отрицательными или равны нулю!")
		} else {
			fmt.Println(intToRoman(result))
		}
	}
	// Вывод результата для арабских чисел
	if counter == 0 && a > 0 && a <= 10 && b > 0 && b <= 10 {
		fmt.Println(result)
	}

}
