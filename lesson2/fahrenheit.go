package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Ведите температуру в градусах по шкале Фаренгейта(°F)")
	var tempFahrenheit float64
	_, err := fmt.Scanln(&tempFahrenheit)
	if err != nil {
		fmt.Println("Некорректный ввод. Поддерживается ввод целых и дробных чисел с точкой в качестве разделителя.", err)
		return
	}
	fmt.Println("Исходные данные: температура по шкале Фаренгейта(°F)", strconv.FormatFloat(tempFahrenheit, 'f', 2, 32))

	tempCelsius := (tempFahrenheit - 32) * 5 / 9
	fmt.Print("Вычисляю температуру по шкале Цельсия")
	for i := 0; i < 7; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("")
	fmt.Println("Температура по шкале Цельсия:", strconv.FormatFloat(tempCelsius, 'f', 2, 32))
}
