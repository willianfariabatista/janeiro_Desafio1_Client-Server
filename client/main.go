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

// Mesma estrutura usada no server para receber o "bid"
type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	// Cria um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Monta a requisição para o endpoint do servidor Go
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição:", err)
	}

	// Executa a requisição
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro na requisição ou timeout excedido: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Recebido status diferente de 200: %d\n", resp.StatusCode)
		return
	}

	// Lê o JSON com o campo "bid"
	var c Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Abre/cria o arquivo cotacao.txt para salvar a cotação
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println("Erro ao criar arquivo cotacao.txt:", err)
		return
	}
	defer file.Close()

	// Escreve no arquivo no formato solicitado: "Dólar: {valor}"
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", c.Bid))
	if err != nil {
		log.Println("Erro ao escrever no arquivo:", err)
		return
	}

	log.Println("Cotação salva em cotacao.txt com sucesso!")
}
