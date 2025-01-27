package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Estrutura para receber o "bid" que será retornado ao cliente
type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/cotacao", handlerCotacao)

	log.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerCotacao(w http.ResponseWriter, r *http.Request) {
	// Contexto da requisição (para poder cancelar caso o cliente encerre antes)
	ctx := r.Context()
	log.Println("Request iniciada.")
	defer log.Println("Request finalizada.")

	// Aqui fazemos a chamada à API de cotação do Dólar
	cotacao, err := buscarCotacao(ctx)
	if err != nil {
		log.Println("Erro ao buscar cotação:", err)
		http.Error(w, "Erro interno ao buscar cotação.", http.StatusInternalServerError)
		return
	}

	// Retorna para o cliente apenas o valor do "bid" em formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cotacao); err != nil {
		log.Println("Erro ao codificar resposta JSON:", err)
		http.Error(w, "Erro ao gerar JSON.", http.StatusInternalServerError)
		return
	}
}

func buscarCotacao(ctx context.Context) (Cotacao, error) {
	// Monta a requisição HTTP com o contexto (permite cancelamento/timeout)
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return Cotacao{}, err
	}

	// Realiza a chamada ao serviço externo
	client := &http.Client{
		// Você pode definir um timeout aqui se desejar limitar o tempo de resposta do servidor externo
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return Cotacao{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Cotacao{}, err
	}

	var dadosAPI map[string]map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&dadosAPI); err != nil {
		return Cotacao{}, err
	}

	// Pega o valor do campo "bid" dentro de "USDBRL"
	bid := dadosAPI["USDBRL"]["bid"]

	return Cotacao{Bid: bid}, nil
}
