package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Estrutura que representa a cotação do dólar.
type Cotacao struct {

	// Armazena o valor da cotação e é mapeado para a chave "bid" do JSON.
	Bid string `json:"bid"`
}

func main() {

	// Cria um contexto com timeout de 300ms.
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Monta a requisição para o endpoint do servidor Go Local.
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição:", err)
	}

	// Envia a requisição HTTP.
	resp, err := http.DefaultClient.Do(req)

	// Trata erros de requisição ou timeout.
	if err != nil {
		log.Printf("Erro na requisição ou timeout excedido: %v\n", err)
		return
	}

	// Fecha o corpo da resposta HTTP (Execupa por ultimo utilizando o "defer") para evitar vazamento de recurso.
	defer resp.Body.Close()

	// Verifica se a resposta HTTP tem status diferente de 200 (OK).
	if resp.StatusCode != http.StatusOK {
		log.Printf("Recebido status diferente de 200: %d\n", resp.StatusCode)
		return
	}

	var c Cotacao

	// Decodifica a resposta JSON do corpo da requisição HTTP para a estrutura Cotacao.
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Cria um arquivo chamado "cotacao.txt" para armazenar os dados.
	file, err := os.Create("cotacao.txt")

	// Se ocorrer um erro na criação do arquivo registra no log.
	if err != nil {
		log.Println("Erro ao criar arquivo cotacao.txt:", err)
		return
	}

	// Garante que o arquivo seja fechado.
	defer file.Close()

	// Escreve a cotação do dólar no arquivo "cotacao.txt".
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", c.Bid))

	// Se ocorrer um erro na escrita registra no log
	if err != nil {
		log.Println("Erro ao escrever no arquivo:", err)
		return
	}

	// Registra no log que a cotação foi salva com sucesso no arquivo "cotacao.txt".
	log.Println("Cotação salva em cotacao.txt com sucesso!")
}
