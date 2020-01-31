package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Ведите длину первого катета")
	var cathet1 float64
	_, err := fmt.Scanln(&cathet1)
	if err != nil {
		fmt.Println("Некорректный ввод. Поддерживается ввод целых и дробных чисел с точкой в качестве разделителя.", err)
		return
	}

	fmt.Println("Ведите длину второго катета")
	var cathet2 float64
	_, err = fmt.Scanln(&cathet2)
	if err != nil {
		fmt.Println("Некорректный ввод. Поддерживается ввод целых и дробных чисел с точкой в качестве разделителя.", err)
		return
	}

	hypotenuse := math.Hypot(cathet1, cathet2)
	fmt.Println("Длина гипотенузы:", strconv.FormatFloat(hypotenuse, 'f', 2, 32))
}
