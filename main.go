package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64 // рост
	var userKg float64     // вес
	fmt.Print("⚖️___ Калькулятор индекса массы тела ___ ⚖️\n")
	fmt.Print("Enter your user height (в метрах): ")
	_, err := fmt.Scan(&userHeight)
	if err != nil {
		fmt.Printf("Вы неправильно ввели ваш рост. Например: 1.92 ERROR: %s", err)
		return
	}
	fmt.Println("Enter your user kg: ")
	_, err = fmt.Scan(&userKg)
	if err != nil {
		fmt.Printf("Вы неправильно ввели ваш вес. Например: 90 ERROR: %s", err)
	}
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Printf("Ваш индекс массы тела: %f", IMT)
}
