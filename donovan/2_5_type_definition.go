package main

import "fmt"

// Определение типа
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

var initValue string

func init() {
	initValue = "Это значение установится при сборке пакета"
}

func init() {
	initValue = initValue + ". Функций init может быть множество"
}

func main() {
	fmt.Println(celsiusToFahrenheit(FreezingC))

	// Так нельзя - разные типы!
	// fmt.Println(Celsius(0) == Fahrenheit(32))

	// Но можно сравнить с базовым
	fmt.Println(Celsius(0) == 32)

	// Или использовать приведение типа
	fmt.Println(Celsius(0) == Celsius(Fahrenheit(32)))

	fmt.Println(AbsoluteZeroC.String())

	fmt.Println(initValue)
}

func celsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // Создание значения нового типа Fahrenheit()
}

func fahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Связываем функцию с типом
func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}
