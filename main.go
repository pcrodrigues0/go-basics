package main

// ** Const declarations **
// const nome string = "paulo"
// const nome = "Paulo"

// type Cliente struct {
// 	Nome     string
// 	Idade    int
// 	Endereco Endereco
// 	Email    string
// }

// type Endereco struct {
// 	Rua    string
// 	Numero int
// 	Bairro string
// }

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

	// data objects
	// var gavetas [2]string
	// gavetas[0] = "Copos"
	// gavetas[1] = "Prato"
	// // gavetas[2] = "123" O array é imutavel
	// fmt.Println(gavetas[0])

	// var gavetas []string // isso é um slice
	// gavetas = append(gavetas, "Copo", "prato", "talher")
	// // fmt.Println(len(gavetas)) // Verificar tamanho do slice
	// // fmt.Println(gavetas[1:2]) // Slice[x:x-1] (aqui é melhor imaginar que o primeiro parametro se trata do ponteiro [0,1,2] e o segundo parametro está atrlado ao len [1,2,3])
	// // reatribuindo valor no slice baseado em um range dele mesmo
	// gavetas = gavetas[:2] // Inserindo até o length 2
	// fmt.Println(gavetas)

	// Maps key -> value
	// var pessoas = map[string]int{} // Aqui foi definido que a chave é uma string e o valor é um int
	// pessoas["paulo"] = 30          // Aqui a chave é unica
	// pessoas["paulo1"] = 32
	// if idade, ok := pessoas["paulo"]; ok {
	// 	fmt.Println("Existe no map", idade, ok)
	// } else {
	// 	fmt.Println("não existe no map")
	// }
	// delete(pessoas, "paulo1")
	// fmt.Println(pessoas)

	// Estruturas de controle

	// nota := 75

	// if nota >= 90 {
	// 	fmt.Println("Aprovado")
	// } else if nota >= 70 {
	// 	fmt.Println("fez o basico")
	// } else {
	// 	fmt.Println("Salsifufu")
	// }

	// if err := thisIsAnError(); err != nil { // Atribuindo o valor da função direto pra uma variavel dentro da propria verificação! Essa variável está dentro do bloco do if, em todo o restante ela é undefined
	// 	fmt.Println(err.Error())
	// }

	// players := map[string]int{
	// 	"paulo": 30,
	// }
	// if pontos, ok := players["paulo"]; ok { // Acessar a chave de um map retorna dois valores, um é o valor de fato e o segundo é se ele existe
	// 	fmt.Println(pontos)
	// }

	// Switch case
	// fmt.Println("Quando é terça?")
	// today := time.Now().Weekday()

	// switch time.Tuesday {
	// case today + 0:
	// 	fmt.Println("é hoje")
	// case today + 1:
	// 	fmt.Println("é amanha")
	// case today + 2:
	// 	fmt.Println("depois de amanha")
	// default:
	// 	fmt.Println("ta longe")
	// }

	// fooooor!!
	// for inicializacao; expressao; fim interacao {

	// }
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println("result", sum) //45

	// Simulando o while
	// sum := 0
	// for sum < 20 {
	// 	fmt.Println("loop")
	// 	sum += 1
	// }
	// fmt.Println(sum)

	// fazendo for baseado em slice
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// range
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// for key, value := range nums {
	// 	fmt.Println(key, value)

	// }
	// range usando map
	// users := map[string]string{
	// 	"nome":      "Paulo",
	// 	"sobrenome": "teste",
	// }
	// for key, value := range users {
	// 	fmt.Println(key, value)
	// }
	// for _, value := range users {
	// 	fmt.Println(value)
	// }
	// for key, _ := range users {
	// 	fmt.Println(key)
	// }

	// Conceito de structs (Viva o typescript mais simples)

	// cliente1 := Cliente{
	// 	Nome:  "Paulo",
	// 	Idade: 29,
	// 	Endereco: Endereco{
	// 		Rua: "227",
	// 	},
	// }
	// cliente1.Email = "emailteste"
	// cliente1.Endereco.Numero = 12
	// fmt.Println(cliente1)

}

// func thisIsAnError() error {
// 	return errors.New("erro generico")
// }
