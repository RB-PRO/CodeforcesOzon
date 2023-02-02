package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() { // Главная функция
	inputBuf := bufio.NewReader(os.Stdin)   // Считывание
	outputBuf := bufio.NewWriter(os.Stdout) // Вывод
	fmt.Fprintln(outputBuf, Work(inputBuf)) // Записываем в вывод
	outputBuf.Flush()                       // Опустошаем вывод
}

// Решение
func Work(in *bufio.Reader) (outputStr string) {
	var testCount int

	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var n, dataInt int
		fmt.Fscan(in, &n) // Прочитать к-во товаров

		var arrInt []int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &dataInt) // Прочитать позиции товаров
			arrInt = append(arrInt, dataInt)
		}

		// Сохранение результата
		outputStr += strconv.Itoa(FindCoast(arrInt)) + "\n"
	}

	return outputStr
}

// Поиск цены по массиву товаров
func FindCoast(inp []int) int {
	// Идея - создать мапу[цена]количество,
	// далее обработать её, рассчитав необходимые значения.

	coasts := make(map[int]int)   // Выделяем память под мапу
	for _, valCost := range inp { // Цикл по входному массиву
		coasts[valCost]++ // Заполнение цены
	}

	// Итак. Если остаток от деления к-ва товаров на 3 равен 0, т.е. мы купили 3 товара, то пишем двойную цену. В противном случае ничего не меняется

	var resultCoast int
	for key, val := range coasts {

		// Если товаторов мало
		if val < 3 {
			resultCoast += val * key
			continue
		}

		// Если к-во товаров кратно 3
		if val%3 == 0 { // Если остаток от деления к-ва товаров равно нулю
			resultCoast += (val / 3 * 2) * key // То записываю 2 из 3 товаров в цену
		} else { // В противном случае остаёются товары, которых больше  или равно 4
			resultCoast += ((val-val%3)/3*2)*key + (val%3)*key
		}

	}

	return resultCoast
}
