package main

//Resumo
//Crie o módulo: go mod init <algum-nome>
//Adicione a dependência: go get github.com/mattn/go-sqlite3
//Compile ou rode seu projeto normalmente: go build, go run, etc.
//Caso você já tenha um arquivo go.mod, basta fazer o passo 2 (o go get). Se não tiver, é fundamental criar primeiro (passo 1).

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	//_ "github.com/mattn/go-sqlite3" // Driver do SQLite

	_ "modernc.org/sqlite" // Driver SQLite
)

// Estrutura que representa a cotação do dólar.
type Cotacao struct {
	// Armazena o valor da cotação e é mapeado para a chave "bid" do JSON.
	Bid string `json:"bid"`
}

func main() {
	// Definindo um endpoint /cotacao.
	http.HandleFunc("/cotacao", handlerCotacao)

	// Mensagem indicando que o Servidor foi iniciado.
	log.Println("Servidor iniciado na porta :8080")

	// Caso ocorra erro ao iniciar o servidor, registra no log e encerra.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerCotacao(w http.ResponseWriter, r *http.Request) {
	// Contexto básico da requisição.
	ctx := r.Context()

	log.Println("Request iniciada.")
	defer log.Println("Request finalizada.")

	// Obter a cotação do dólar com timeout de 200ms
	ctxAPICall, cancelAPICall := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancelAPICall()

	cotacao, err := buscarCotacao(ctxAPICall)
	if err != nil {
		log.Println("Erro ao buscar cotação:", err)
		http.Error(w, "Erro interno ao buscar cotação.", http.StatusInternalServerError)
		return
	}

	// Persistir a cotação no banco com timeout de 10ms
	ctxDB, cancelDB := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancelDB()

	if err := salvarCotacaoNoBanco(ctxDB, cotacao); err != nil {
		log.Println("Erro ao salvar cotação no banco:", err)
		http.Error(w, "Erro interno ao salvar cotação.", http.StatusInternalServerError)
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

// buscarCotacao realiza a chamada externa para obter a cotação do dólar.
func buscarCotacao(ctx context.Context) (Cotacao, error) {
	// Monta a requisição HTTP com o contexto (permite cancelamento/timeout).
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return Cotacao{}, err
	}

	// Define um client com o timeout de 200 ms (ou menos, se o contexto expirar antes).
	client := &http.Client{
		Timeout: 200 * time.Millisecond,
	}

	// Executa a requisição.
	resp, err := client.Do(req)
	if err != nil {
		return Cotacao{}, err
	}
	defer resp.Body.Close()

	// Verifica status da resposta.
	if resp.StatusCode != http.StatusOK {
		return Cotacao{}, fmt.Errorf("status inesperado: %d", resp.StatusCode)
	}

	// Decodifica a resposta em um mapa para extrair apenas o campo "bid".
	var dadosAPI map[string]map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&dadosAPI); err != nil {
		return Cotacao{}, err
	}

	// Extrai o "bid" do mapa.
	bid := dadosAPI["USDBRL"]["bid"]

	return Cotacao{Bid: bid}, nil
}

// salvarCotacaoNoBanco grava a cotação no banco SQLite.
func salvarCotacaoNoBanco(ctx context.Context, c Cotacao) error {
	// Abre (ou cria) o arquivo/banco "cotacoes.db".
	// db, err := sql.Open("sqlite3", "cotacoes.db")
	db, err := sql.Open("sqlite", "cotacoes.db") // para modernc.org/sqlite
	if err != nil {
		return fmt.Errorf("falha ao abrir o banco: %w", err)
	}
	defer db.Close()

	// Cria a tabela caso não exista (id, bid, data/hora).
	createTable := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err = db.ExecContext(ctx, createTable); err != nil {
		return fmt.Errorf("falha ao criar tabela: %w", err)
	}

	// Insere a cotação.
	insertStmt := `INSERT INTO cotacoes (bid) VALUES (?)`
	if _, err = db.ExecContext(ctx, insertStmt, c.Bid); err != nil {
		return fmt.Errorf("falha ao inserir cotacao: %w", err)
	}

	log.Println("Cotação salva no banco de dados SQLite com sucesso!")
	return nil
}
