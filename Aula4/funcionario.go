package main

import "fmt"

type Funcionario struct {
	nome    string
	idade   int
	funcao  string
	salario int
}

func main() {
	//cria array funcionario
	funcionarios := []Funcionario{
		{"Julia", 21, "estagiaria", 2000},
		{"Thomas", 20, "estagiario", 2000},
		{"Maria", 15, "sem funcao", 0},
	}
	//pesquisa pela posição 1 do array (funcionario 2)
	fmt.Println("Funcionario 2: ", funcionarios[1])
}
