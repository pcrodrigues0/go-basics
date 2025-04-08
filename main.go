package main

import (
	"fmt"
	"strings"
)

// ** Const declarations **
// const nome string = "paulo"
// const nome = "Paulo"

func main() {
	// var texto string = "Ola"
	// texto = "tchau"
	// var texto = "Declaracao direta"
	// ** Declaração de valor e somando **
	// var texto = 123
	// texto += 5
	// var texto string
	// var texto int
	// var texto bool
	// ** Declaracao curta **
	// texto := "Ola"
	// texto := 123
	// texto -= 1
	// fmt.Println(texto)

	// Constants
	// nome = "123"
	// fmt.Println(nome)

	// Forcando tamanho do int
	// - **int8**: Inteiro de 8 bits (valores de -128 a 127)
	// -**int16**: Inteiro de 16 bits (valores de -32768 a 32767)
	// -**int32**: Inteiro de 32 bits (valores de -2147483648 a 2147483647)
	// -**int64**: Inteiro de 64 bits (valores de -9223372036854775808 a 9223372036854775807)
	// var idade int = 30
	// var contador int32 = 2
	// var indice int8 = 1

	// fmt.Println("Idade:", idade)
	// fmt.Println("Idade:", contador)
	// fmt.Println("Idade:", indice)

	// Numeros de ponto flutuante
	// Não e possivel realizar operações de tipos diferentes de float
	// var floatNumber float32 = 1.1
	// var pi float64 = 3.14
	// var raio float64 = 2.5
	// var area = pi * raio * raio
	// fmt.Println("FloatNumber:", floatNumber)
	// fmt.Println("Area do circulo:", area)

	// Tipo booleano
	// var novotipo bool
	// var maior bool = true
	// var maiorCalc bool = 10 > 5
	// fmt.Println(novotipo)
	// fmt.Println(maior)
	// fmt.Println(maiorCalc)

	// string handling
	// var hello string = "Olá Você"
	// var greet string = " Tudo certo?"
	// var meet = hello + greet
	// fmt.Println(meet)
	// fmt.Println(strings.ToUpper(meet))
	// fmt.Println(strings.Contains(meet, "tudo"))    // Essa merda acaba levando em consideração case-sensitive
	// fmt.Println(strings.Contains(meet, "Tudo"))    // Assim buscou
	// fmt.Println(strings.ContainsAny(meet, "tudo")) //ContainsAny busca por unicode...

}
