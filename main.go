package main

import (
	"fmt"
	"math"
	"errors"
)

const IMTPower = 2

func main() {

	fmt.Println("⚖️___ Калькулятор индекса массы тела ___ ⚖️")
	for {
		userHeight, userKg := getUserInput()
		IMT, err := calculateIMT(userHeight, userKg)
		if err != nil {
			fmt.Println("Неправильно заданы параметры рассчёта:", err)
			continue
			// panic("Неправильно заданы параметры рассчёта")
		}
		outputResult(IMT)
		isRepeateCalculation := checkRepeateCalculation()
		if !isRepeateCalculation{
			break
		}
	}

}

func outputResult(imt float64) {
	fmt.Printf("✅ Ваш индекс массы тела: %.0f\n", imt)
	
	switch {
	case imt < 16:
		fmt.Println("🦴 У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("🍖 У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("😎 У вас нормальный вес")
	case imt < 30:
		fmt.Println("👀 У вас избыточный вес")
	case imt < 40:
		fmt.Println("😱 У вас 2-я степень ожирения")
	case imt >= 40:
		fmt.Println("🤯 У вас 3-я степень ожирения")
	default:
		fmt.Println("Не можем определить ваш вес")
	}
}

func calculateIMT(userHeight, userKg float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("USER_INPUT_ERROR")
	}
	result := userKg / math.Pow(userHeight/100, IMTPower)
	return result, nil
}

func checkRepeateCalculation() bool{
	var userChoise string
	fmt.Print("Вы хотите повторить работу приложения? Y/N: ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64

	fmt.Print("Введите ваш рост (в сантиметрах, например: 192.1): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес (в кг, например: 90): ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}