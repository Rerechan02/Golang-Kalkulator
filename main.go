package main

import (
        "fmt"
        "os"
        "os/exec"
)

func clearScreen() {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}

func showHeader() {
        fmt.Println("====================================")
        fmt.Println("           KALKULATOR")
        fmt.Println("====================================")
}

func showMenu() {
        fmt.Println("Pilih operasi yang ingin dilakukan:")
        fmt.Println("1. Penjumlahan")
        fmt.Println("2. Pengurangan")
        fmt.Println("3. Perkalian")
        fmt.Println("4. Pembagian")
        fmt.Println("5. Keluar")
}

func add(num1, num2 float64) float64 {
        return num1 + num2
}

func subtract(num1, num2 float64) float64 {
        return num1 - num2
}

func multiply(num1, num2 float64) float64 {
        return num1 * num2
}

func divide(num1, num2 float64) float64 {
        return num1 / num2
}

func main() {
        clearScreen()
        showHeader()

        for {
                showMenu()
                var choice string
                fmt.Print("Pilih operasi (1-5): ")
                fmt.Scanln(&choice)

                if choice == "5" {
                        fmt.Println("Terima kasih telah menggunakan kalkulator!")
                        break
                }

                clearScreen()
                showHeader()

                var num1, num2 float64
                fmt.Print("Masukkan angka pertama: ")
                fmt.Scanln(&num1)
                fmt.Print("Masukkan angka kedua: ")
                fmt.Scanln(&num2)

                var hasil float64
                switch choice {
                case "1":
                        hasil = add(num1, num2)
                case "2":
                        hasil = subtract(num1, num2)
                case "3":
                        hasil = multiply(num1, num2)
                case "4":
                        hasil = divide(num1, num2)
                default:
                        fmt.Println("Pilihan tidak valid")
                        continue
                }

                clearScreen()
                showHeader()
                fmt.Printf("Angka pertama: %.2f\n", num1)
                fmt.Printf("Angka kedua: %.2f\n", num2)
                fmt.Printf("Hasil: %.2f\n", hasil)
                fmt.Println()
                fmt.Print("Tekan Enter untuk melanjutkan...")
                fmt.Scanln()
                clearScreen()
        }
