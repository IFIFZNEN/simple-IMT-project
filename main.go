package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64 // рост
	var userKg float64     // вес
	fmt.Println("⚖️___ Калькулятор индекса массы тела ___ ⚖️")
	fmt.Print("Enter your user height (в сантиметрах): ")
	_, err := fmt.Scan(&userHeight)
	if err != nil {
		fmt.Printf("Вы неправильно ввели ваш рост. Например: 1.92 ERROR: %s", err)
		return
	}
	fmt.Print("Enter your user kg: ")
	_, err = fmt.Scan(&userKg)
	if err != nil {
		fmt.Printf("Вы неправильно ввели ваш вес. Например: 90 ERROR: %s", err)
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %0.f", imt)
	fmt.Print(result)
}
