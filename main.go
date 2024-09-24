package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Livro struct {
	Nome  string `json:"nome"`
	Autor string `json:"autor"`
	Ano   int    `json:"ano"`
}

func main() {

	// ler arquivo json
	arquivo, err := os.Open("ATD/01_dados.json")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %s", err)
	}
	defer arquivo.Close()

	// ler conteudo do arquivo
	conteudo, err := ioutil.ReadAll(arquivo)
	if err != nil {
		log.Fatalf("Erro ao ler o conteudo do arquivo: %s", err)
	}

	// Deserializar (decodificar) o JSON para um slice de struct
	var livros []Livro
	err = json.Unmarshal(conteudo, &livros)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %s", err)
	}

	// Add um novo livro
	novoLivro := Livro{
		Nome:  "Psicologia Comportamental",
		Autor: "Autor C",
		Ano:   2022,
	}

	nuevoLibro := Livro{
		Nome:  "Psicologia do Cu",
		Autor: "C de CU",
		Ano:   2024,
	}
	livros = append(livros, novoLivro)
	livros = append(livros, nuevoLibro)

	// Exibir dados alterados
	for _, livro := range livros {
		fmt.Printf("%s por %s (%d)\n", livro.Nome, livro.Autor, livro.Ano)
	}

	// Serializar (codificar) o slice de struct de volta para JSON
	dadosModificados, err := json.MarshalIndent(livros, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao serializar o JSON: %s", err)
	}

	// Escrever dados modificados no arquivo
	err = ioutil.WriteFile("ATD/01_dados_modificados.json", dadosModificados, 0644)
	if err != nil {
		log.Fatalf("Erro ao escrever no arquivo: %s", err)
	}
}
