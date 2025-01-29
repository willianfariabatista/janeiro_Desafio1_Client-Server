package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Estrutura que representa a cotação do dólar.
type Cotacao struct {

	// Armazena o valor da cotação e é mapeado para a chave "bid" do JSON.
	Bid string `json:"bid"`
}

func main() {

	// Definindo um endpoint.
	http.HandleFunc("/cotacao", handlerCotacao)

	// Mensagem indicando que o Servidor foi iniciado.
	log.Println("Servidor iniciado na porta :8080")

	// Caso Ocorra erro ao iniciar o servidor retorna no Log.Fatal.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerCotacao(w http.ResponseWriter, r *http.Request) {

	// Contexto da requisição (para poder cancelar caso o cliente encerre antes).
	ctx := r.Context()
	log.Println("Request iniciada.")
	defer log.Println("Request finalizada.")

	// Fazendo a chamada a API de cotação do Dólar.
	cotacao, err := buscarCotacao(ctx)
	if err != nil {
		log.Println("Erro ao buscar cotação:", err)
		http.Error(w, "Erro interno ao buscar cotação.", http.StatusInternalServerError)
		return
	}

	// Retorna para o cliente apenas o valor do "bid" em formato JSON.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cotacao); err != nil {
		log.Println("Erro ao codificar resposta JSON:", err)
		http.Error(w, "Erro ao gerar JSON.", http.StatusInternalServerError)
		return
	}
}

func buscarCotacao(ctx context.Context) (Cotacao, error) {

	// Monta a requisição HTTP com o contexto (permite cancelamento/timeout).
	// Se ocorrer um erro na criação da requisição, retorna uma estrutura vazia e o erro.
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return Cotacao{}, err
	}

	// Realiza a chamada ao serviço externo.
	client := &http.Client{

		// tempo limite de 2 segundos para a requisição.
		Timeout: 2 * time.Second,
	}

	// Envia a requisição HTTP para obter a cotação do dólar.
	resp, err := client.Do(req)

	// Se ocorrer um erro na requisição, retorna uma estrutura vazia e o erro.
	if err != nil {
		return Cotacao{}, err
	}

	// Garante que o corpo da resposta HTTP seja fechado.
	defer resp.Body.Close()

	// Verifica se a resposta HTTP tem status diferente de 200 (OK). Se for diferente, retorna uma estrutura vazia e o erro.
	if resp.StatusCode != http.StatusOK {
		return Cotacao{}, err
	}

	// Decodifica a resposta JSON da API em um mapa aninhado de strings.
	var dadosAPI map[string]map[string]string

	// Se ocorrer um erro na decodificação, retorna uma estrutura vazia e o erro.
	if err := json.NewDecoder(resp.Body).Decode(&dadosAPI); err != nil {
		return Cotacao{}, err
	}

	// Obtem o valor da cotação do dólar (campo "bid") a partir do mapa de resposta da API.
	bid := dadosAPI["USDBRL"]["bid"]

	// Retorna uma estrutura Cotacao com o valor obtido da API
	return Cotacao{Bid: bid}, nil
}
