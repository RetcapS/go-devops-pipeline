package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("--- Go Hesap Makinesi ---")

	// Kullanıcıdan ilk sayıyı alma
	fmt.Print("İlk sayıyı girin: ")
	fmt.Scan(&num1)

	// Kullanıcıdan işlemi alma
	fmt.Print("İşlemi girin (+, -, *, /): ")
	fmt.Scan(&operator)

	// Kullanıcıdan ikinci sayıyı alma
	fmt.Print("İkinci sayıyı girin: ")
	fmt.Scan(&num2)

	var result float64

	// Girilen operatöre göre işlemi yapma
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Hata: Bir sayı sıfıra bölünemez!")
			return // Programı güvenli bir şekilde sonlandır
		}
		result = num1 / num2
	default:
		fmt.Println("Hata: Geçersiz bir işlem (+, -, *, /) girdiniz!")
		return
	}

	// Sonucu ekrana yazdırma
	fmt.Printf("\nSonuç: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
