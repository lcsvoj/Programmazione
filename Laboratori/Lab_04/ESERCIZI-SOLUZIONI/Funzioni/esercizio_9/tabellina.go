package main

import "fmt"

func main() {

	var numero int

	for {
		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&numero)

		if !Tabellina(numero) {
			break
		}
	}

	fmt.Println("Programma terminato.")

}

func Tabellina(numero int) bool {
	if numero >= 1 && numero <= 9 {
		fmt.Print("Tabellina del ", numero, ":")
		var contatore int
		for contatore = 0; contatore <= 10; contatore++ {
			fmt.Print(" ", contatore*numero)
		}
		fmt.Println()
		return true
	}
	return false
}
