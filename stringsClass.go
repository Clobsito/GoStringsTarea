package main

import (
	"fmt"
	"strings"
)

func main() {

	var opcion int = 0
	var cadena, nombre, palabra, caracter string = "", "", "", ""
	var resp bool

	fmt.Print("Introduzca su nombre: ");
	fmt.Scan(&nombre);

	fmt.Println("\nBienvenido al sistema,", nombre);

	fmt.Println("\nQué operación desea realizar?");
	fmt.Println("\n1. Aplicar mayusculas a una cadena");
	fmt.Println("2. Aplicar minusculas a una cadena");
	fmt.Println("3. Contar las veces que un caracter aparecio en una cadena");
	fmt.Println("4. Buscar caracteres en una cadena");
	fmt.Println("5. Determinar longitud de una cadena");
	fmt.Println("");
	fmt.Scan(&opcion)

	if (opcion == 1){

		fmt.Print("\nIntroduzca una cadena de texto en minusculas: ");
		fmt.Scan(&cadena);

		fmt.Print("Su cadena ", "'",cadena,"'", " se convirtio en: ","'",strings.ToUpper(cadena),"'");

	} else if(opcion == 2){

		fmt.Print("\nIntroduzca una cadena de texto en mayusculas: \n");
		fmt.Scan(&cadena);

		fmt.Print("Su cadena ", "'",cadena,"'", " se convirtio en: ", "'",strings.ToLower(cadena),"'");

	} else if(opcion == 3){

		fmt.Print("\nIntroduzca una cadena de texto: ");
		fmt.Scan(&cadena);

		fmt.Print("\nIntroduzca el caracter que desea contar: ")
		fmt.Scan(&caracter);

		fmt.Print("\nEl caracter ","'", caracter, "'", " se encontró: ", strings.Count(cadena, caracter), " vez/veces")
		
	} else if(opcion == 4){

		fmt.Print("\nIntroduzca una cadena de texto: ");
		fmt.Scan(&cadena);

		fmt.Print("\nIntroduzcca los caracteres que desea encontrar: ");
		fmt.Scan(&palabra);

		resp = strings.Contains(cadena, palabra);

		if(resp == true){
			print("\nEl/Los caracteres: ", "'",palabra, "'", " se encuentran en la cadena");
		} else {
			print("\nEl/Los caracteres ", "'",palabra, "'", " no se encuentran en la cadena ");
		}
		
	} else if(opcion == 5){

		fmt.Print("\nIntroduzca una cadena de texto:");
		fmt.Scan(&cadena);

		fmt.Print("La cadena ", "'",cadena,"'", " tiene: ", len(cadena), " caracteres");
		
	} else {
		fmt.Print("\nOpcion invalida");
	}
}