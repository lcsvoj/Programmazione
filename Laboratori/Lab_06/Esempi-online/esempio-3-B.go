// Lettura di testo tramite Scanner
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	var testo string
	fmt.Println("Inserisci del testo su più righe " +
		"(termina con CTRL+D o riga vuota). ")
		
	// premendo la combinazione di tasti Ctrl+D, viene inserito l'indicatore End-Of-File (EOF);
		
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Inserisci una riga:")
	
	for scanner.Scan() {
		
		testo = scanner.Text()
		//if testo == "" { // blocco da commentare se le righe vuote non devono far terminare la lettura;
		//	break
		//}
		fmt.Print("Letto: ***")
		fmt.Print(testo)
		fmt.Println("+++")
		fmt.Println("Inserisci una nuova riga:")
	}
}
