package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Definir a estrutura de dados do Produto em Go
type Produto struct {
	ID          int     `json:"id"`
	Nome        string  `json:"nome"`
	Descricao   string  `json:"descricao"`
	Preco       float64 `json:"preco"`
}

func main() {
	// Listar todos os produtos
	listarProdutos()

	// Obter detalhes de um produto por ID
	obterDetalhesDoProduto(1)

	// Criar um novo produto
	novoProduto := Produto{
		Nome:      "Novo Produto",
		Descricao: "Descrição do Novo Produto",
		Preco:     39.99,
	}
	criarProduto(novoProduto)

	// Atualizar um produto existente (Substituir o ID pelo ID do produto real)
	produtoAtualizado := Produto{
		ID:        1,
		Nome:      "Produto Atualizado",
		Descricao: "Descrição do Produto Atualizado",
		Preco:     49.99,
	}
	atualizarProduto(produtoAtualizado)

	// Excluir um produto por ID (Substituir o ID pelo ID do produto real)
	excluirProduto(2)
}

func listarProdutos() {
	resp, err := http.Get("http://localhost:5000/produtos")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lista de Produtos:")
	fmt.Println(string(body))
}

func obterDetalhesDoProduto(id int) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:5000/produtos/%d", id))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Detalhes do Produto:")
	fmt.Println(string(body))
}

func criarProduto(produto Produto) {
	produtoJSON, _ := json.Marshal(produto)
	resp, err := http.Post("http://localhost:5000/produtos", "application/json", bytes.NewBuffer(produtoJSON))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Novo Produto criado com sucesso!")
}

func atualizarProduto(produto Produto) {
	produtoJSON, _ := json.Marshal(produto)
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:5000/produtos/%d", produto.ID), bytes.NewBuffer(produtoJSON))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Produto Atualizado com sucesso!")
}

func excluirProduto(id int) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:5000/produtos/%d", id), nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Produto Excluído com sucesso!")
}
