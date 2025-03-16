package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const IMTPower = 2

func main() {
	fmt.Println("⚖️___ Калькулятор индекса массы тела ___ ⚖️")
	userHeight := getValidInput("Введите ваш рост (в сантиметрах, например: 192.1): ")
	userKg := getValidInput("Введите ваш вес (в кг, например: 90): ")
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)

	switch {
	case IMT < 16:
		fmt.Println("🦴 У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("🍖 У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("😎 У вас нормальный вес")
	case IMT < 30:
		fmt.Println("👀 У вас избыточный вес")
	case IMT < 40:
		fmt.Println("😱 У вас 2-я степень ожирения")
	case IMT >= 40:
		fmt.Println("🤯 У вас 3-я степень ожирения")
	default:
		fmt.Println("Не можем определить ваш вес")
	}
}

func getValidInput(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("❌ Ошибка: Введите корректное число!")
	}
}

func outputResult(imt float64) {
	fmt.Printf("✅ Ваш индекс массы тела: %.0f\n", imt)
}

func calculateIMT(userHeight, userKg float64) float64 {
	result := userKg / math.Pow(userHeight/100, IMTPower)
	return result
}
