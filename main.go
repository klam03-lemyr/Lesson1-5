package main

import "fmt"

func main() {

	fmt.Println(isEven(5))
	fmt.Println(getSign(0))
	fmt.Println(calculate(1, 2, "/"))
	fmt.Println(getSeason(9))
	fmt.Println(triangleType(3, 5, 5))
}

func isEven(number int) string {

	var even string

	if number%2 == 0 {
		even = "Четное"
	} else {
		even = "Нечетное"
	}
	return even
}

func getSign(number int) string {
	var sign string

	if number >= 1 {
		sign = "Положительное"
	} else {
		sign = "Отрицательное"
	}
	if number == 0 {
		sign = "Ноль"
	}

	return sign
}

func calculate(first float64, second float64, operator string) float64 {
	var result float64
	if operator == "+" {
		result = first + second
	} else if operator == "-" {
		result = first - second
	} else if operator == "*" {
		result = first * second
	} else if operator == "/" {
		result = first / second
	} else {
		result = 0
	}
	return result
}

func getSeason(month int) string {
	var season string
	if month == 12 || month == 1 || month == 2 {
		season = "Зима"
	} else if month == 3 || month == 4 || month == 5 {
		season = "Весна"
	} else if month == 6 || month == 7 || month == 8 {
		season = "Лето"
	} else if month == 9 || month == 10 || month == 11 {
		season = "Осень"
	}
	return season
}

func triangleType(a, b, c int) string {
	var Ttype string
	if a == b && b == c {
		Ttype = "Равносторонний"
	} else if a == b || b == c || a == c {
		Ttype = "Равнобедренный"
	} else {
		Ttype = "Разносторонний"
	}
	return Ttype
}
