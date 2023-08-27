package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() { // Главная функция
	inputBuf := bufio.NewReader(os.Stdin)   // Считывание
	outputBuf := bufio.NewWriter(os.Stdout) // Вывод
	fmt.Fprintln(outputBuf, Work(inputBuf)) // Записываем в вывод
	outputBuf.Flush()                       // Опустошаем вывод
}

/*
	четыре однопалубных корабля,
	три двухпалубных корабля,
	два трёхпалубных корабля,
	один четырёхпалубный корабль.
*/

// Решение
func Work(in *bufio.Reader) (outputStr string) {
	var testCount int

	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ { // Цикл по набору данных

		var strData string
		fmt.Fscan(in, &strData)

		outputStr += strings.Join(NumberAvto(strData), " ") + "\n"
	}

	return outputStr
}

// Функция, которая принимает строку символов и на выходе даёт номера, которые могут получиться
func NumberAvto(str string) []string {
	out := make([]string, 0)
	for {
		if len(str) == 0 {
			break
		}
		// fmt.Println(len(str), IsNumberOne(str[:5]), str[:5], IsNumberTwo(str[:4]), str[:4])
		var NumberTwo, NumberOne bool
		if len(str) >= 4 {
			NumberTwo = IsNumberTwo(str[:4])
		}
		if len(str) >= 5 {
			NumberOne = IsNumberOne(str[:5])
		}

		if NumberOne {
			out = append(out, str[:5]+" ")
			str = str[5:]
		}
		if NumberTwo {
			out = append(out, str[:4]+" ")
			str = str[4:]
		}

		if !NumberOne && !NumberTwo {
			return []string{"-"}
		}

	}
	return out
}

// Проверка, что это номер первого вида
//
// автомобильный номер имеет вид буква-цифра-цифра-буква-буква (примеры корректных номеров первого вида: R48FA, O00OO, A99OK);
func IsNumberOne(number string) bool {
	numberRune := []rune(number)
	// fmt.Printf("%#v, %s\n", numberRune, string(numberRune))
	if len(numberRune) == 5 { // Проверка на длину
		//fmt.Println(IsLetter(numberRune[0]), IsFigure(numberRune[1]), IsFigure(numberRune[2]), IsLetter(numberRune[3]), IsLetter(numberRune[4]))
		if IsLetter(numberRune[0]) && IsFigure(numberRune[1]) && IsFigure(numberRune[2]) && IsLetter(numberRune[3]) && IsLetter(numberRune[4]) {
			return true
		}
	}
	return false
}

// Проверка, что это номер второго вида
//
// автомобильный номер имеет вид буква-цифра-буква-буква (примеры корректных номеров второго вида: T7RR, A9PQ, O0OO).
func IsNumberTwo(number string) bool {
	numberRune := []rune(number)
	// fmt.Printf("%#v", numberRune)
	if len(number) == 4 { // Проверка на длину
		if IsLetter(numberRune[0]) && IsFigure(numberRune[1]) && IsLetter(numberRune[2]) && IsLetter(numberRune[3]) {
			return true
		}
	}
	return false
}

// Првоерка что это буква
func IsLetter(r rune) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}
	return true
}

// Првоерка что это цифра
func IsFigure(r rune) bool {
	//fmt.Println(r, '0', '9')
	return r >= '0' && r <= '9'
}
