package main

import (
	"fmt"
	"strconv"
	"strings"
)

// функция rimToArab преобразует римскую форму записи числа в арабскую
func rimToArab(a string) int {
	var sum = 0
	if a == "IV" {
		sum = 4
	} else if a == "IX" {
		sum = 9
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] == 'I' {
				sum += 1
			}
			if a[i] == 'V' {
				sum += 5
			}
			if a[i] == 'X' {
				sum += 10
			}
		}
	}
	return sum
}

// функция arabToRim преобразует арабскую форму записи числа в римскую
func arabToRim(a int) string {
	var count string
	if a == 100 {
		count += "C"
		a -= 100
	}
	if a >= 90 {
		count += "XC"
		a -= 90
	} else if a >= 50 {
		count += "L"
		a -= 50
	} else if a >= 40 {
		count += "XL"
		a -= 40
	}
	for a >= 10 {
		count += "X"
		a -= 10
	}
	if a == 9 {
		count += "IX"
		a -= 9
	} else if a >= 5 {
		count += "V"
		a -= 5
	} else if a >= 4 {
		count += "IV"
		a -= 4
	}
	for a >= 1 {
		count += "I"
		a -= 1
	}
	return count
}

// функция stringToInteger преобразует строковую форму записи переменных в числовую
func stringToInteger(a string, b string) (int, int, bool) {
	var x, y int
	var c bool
	if strings.ContainsAny(a, "IVX") && strings.ContainsAny(b, "IVX") {
		x = rimToArab(a)
		y = rimToArab(b)
		c = true
		// булевая переменная c передает программе информацию о том, что введенные пользователем числа были римскими

	} else if strings.ContainsAny(a, "1234567890") && strings.ContainsAny(b, "1234567890") {
		x, _ = strconv.Atoi(a)
		y, _ = strconv.Atoi(b)
	}
	return x, y, c
}

// функция wholeNumber проверяет, были ли введеные числа с плавающей точкой
func wholeNumber(a string, b string) bool {
	var x, y, g, h float64
	var m, n int
	var d bool
	//конвертируем исходную строковую форму записи в число с плавающей точкой
	x, _ = strconv.ParseFloat(a, 64)
	y, _ = strconv.ParseFloat(b, 64)
	//конвертируем исходную строковую форму записи в целое число
	m, _ = strconv.Atoi(a)
	n, _ = strconv.Atoi(b)
	//переводим целое число в число с плавающей точкой
	g = float64(m)
	h = float64(n)
	//сравниваем два значения, если они равны то исходная запись была целое число
	if g == x && h == y {
		d = true
	} else {
		d = false
	}
	return d
}

// функция rimWithArab проверяет, совпадают ли системы счисления двух слагаемых
func rimWithArab(a string, b string) bool {
	var e bool
	//сравниваем, содержат ли два аргумента римскую и арабскую форму записи одновременно
	if strings.ContainsAny(a, "1234567890") && strings.ContainsAny(b, "IVX") {
		e = true
	} else if strings.ContainsAny(b, "1234567890") && strings.ContainsAny(a, "IVX") {
		e = true
	}
	return e
}

// функция interval проверяет, чтобы введеные числа попадали в диапазон 1<x<=10
func interval(a int, b int) bool {
	var f bool
	if a > 0 && a <= 10 && b > 0 && b <= 10 {
		f = true
	} else {
		f = false
	}
	return f
}

func main() {
	var argument1 string
	var argument2 string
	var operand string
	var result int
	var count int
	_, err := fmt.Scanln(&argument1, &operand, &argument2)
	if err != nil { // проверяем, чтобы было введено 2 аргумента и один операнд
		panic("Ошибка! Введите два оператора и один операнд.")
	}
	argumentInteger1, argumentInteger2, isRomain := stringToInteger(argument1, argument2) // преобразовываем строковый тип аргументов в числовой
	if isRomain == true && operand == "-" {                                               // проверяем, чтобы при вычитании в римской системе не получался отрицательный результат
		fmt.Println("Ошибка! В римской системе нет отрицательных чисел.")
		count += 1
	}
	if wholeNumber(argument1, argument2) == false { // проверяем, чтобы введеное число было целым, в противном случае округляем рациональное число до целого
		fmt.Println("Ошибка! Введите целое число.")
		count += 1
		argumentIntegerFloat1, _ := strconv.ParseFloat(argument1, 64)
		argumentIntegerFloat2, _ := strconv.ParseFloat(argument2, 64)
		argumentInteger1 = int(argumentIntegerFloat1)
		argumentInteger2 = int(argumentIntegerFloat2)
	}
	if interval(argumentInteger1, argumentInteger2) == false { // проверяем, чтобы введеное число было больше 1 и меньше 10
		fmt.Println("Ошибка! Введенное число должно быть от 1 до 10.")
		count += 1
	}
	if rimWithArab(argument1, argument2) == true { //проверяем, чтобы введеные аргументы принадлежали одной системе счисления
		fmt.Println("Ошибка! Используются разные системы счисления одновременно.")
		count += 1
	}
	if count == 0 { // если счетчик ошибок = 0 производим вычисление и выводим результат
		switch operand {
		case "*":
			result = argumentInteger1 * argumentInteger2
		case "/":
			result = argumentInteger1 / argumentInteger2
		case "+":
			result = argumentInteger1 + argumentInteger2
		case "-":
			result = argumentInteger1 - argumentInteger2
		}
		if count == 0 && isRomain == true { // если счетчик ошибок = 0 и аргументы были введены в арабской системе, производим вычисление и выводим результат
			arabToRim(result)
		}
		fmt.Println(result)
	}
}
