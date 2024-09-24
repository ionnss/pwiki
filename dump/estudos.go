package main

/*func main() {


	// concatenação de strings
	fmt.Println("Go" + "Lang")

	// operadores aritmeticos
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3, 0)

	// operadores booleanos
	fmt.Println(true && false) // AND
	fmt.Println(true || false) // OR
	fmt.Println(!true)         // NOT
	fmt.Println(!false)        // NOT

	// variaveis
	var a = "início" // declarando uma variável
	fmt.Println(a)

	var b, c int = 1, 2 // declarando mais de uma variável
	fmt.Println(b, c)

	var d = true // inferindo o tipo da variável
	fmt.Println(d)

	var e int // declarando uma variável sem valor ( zero value )
	fmt.Println(e)

	f := "Mundão" // inferindo o tipo da variável e atribuindo um valor
	fmt.Println(f)

	// constantes
	const s string = "constante" // declarando uma constante
	fmt.Println(s)               // exibindo o valor de uma constante

	const n = 500000000 // declarando uma constante e inferindo o tipo

	const z = 3e20 / n // declarando uma constante e inferindo o tipo
	fmt.Println(z)     // exibindo o valor de uma constante

	fmt.Println(int64(z)) // cast para int64

	fmt.Println(math.Sin(n)) // exibindo o valor de uma constante


/*
	In this example, the loop is controlled by a conditional statement i <= 10.
	The loop will continue to execute as long as the condition is true.
	The variable i is incremented manually inside the loop using i = i + 1.
*/

/*
	i := 1        // declarando e inicializando uma variável
	for i <= 10 { // loop
		fmt.Println(i) // exibindo o valor da variável
		i = i + 1      // incrementando o valor da variável no loop em cada iteração
	}
*/

/*
	In this example, the loop is controlled by a more traditional for loop syntax, which includes:

	- Initialization: j := 0 (sets the initial value of j)
	- Condition: j < 10 (the loop will continue as long as j is less than 10)
	- Increment: j++ (increments j by 1 after each iteration)

	The main difference is that in the second example, the loop control is more concise and expressive,
	with the initialization, condition, and increment all specified in a single line.
	This makes the code easier to read and understand.

	In terms of functionality, both loops will produce the same output:
	printing numbers from 1 to 10 (in the first example) and from 0 to 9 (in the second example).
*/

/*for j := 0; j < 10; j++ { // loop com condicional
		fmt.Println(j)
	}


	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// exemplo para compreender o que é nil em Go
	var p *int     // para armazenar um ponteiro para um inteiro
	fmt.Println(p) // imprime o valor do ponteiro
}
*/
