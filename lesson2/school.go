package main

import (
	"fmt"
)

const studyRoomsCount int = 3
const studentsOnDesk int = 2

func main() {

	var studyCountMap [studyRoomsCount]int
	for i := 0; i < studyRoomsCount; i++ {
		fmt.Println("Укажите количество учеников в классе №", i+1)
		var studyCount int
		_, err := fmt.Scanln(&studyCount)
		if err != nil {
			fmt.Println("Некорректный ввод. Поддерживается только ввод целых чисел.", err)
			return
		}
		studyCountMap[i] = studyCount
	}

	TotalDescCount := 0
	fmt.Println("Исходные данные:")
	for i := 0; i < studyRoomsCount; i++ {
		fmt.Print("Количество учеников в классе №", i+1, " - ", studyCountMap[i])
		tail := int32(studyCountMap[i]) % int32(studentsOnDesk)
		descCount := studyCountMap[i] / studentsOnDesk

		if tail > 0 {
			descCount++
		}
		fmt.Println(". Количество парт к покупке - ", descCount)

		TotalDescCount += descCount
	}

	fmt.Println("Нужно закупить", TotalDescCount, "парт.")
}
