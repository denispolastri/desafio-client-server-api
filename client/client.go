package client

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/denispolastri/desafio-client-server-api/server"
)

func Client() {

	// Logger default
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("Iniciando Client...")
	LeDolarBancoDeDados()
	slog.Info("Finalizando Client...")
}

type DollarBR struct {
	USDBRL server.Dollar `json:"USDBRL"`
}

func LeDolarBancoDeDados() {

	start := time.Now()

	// Cria um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Cria a requisição com o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		slog.Error("erro ao criar a requisição", "error", err)
		return
	}

	// Executa a requisição
	client := &http.Client{}
	request, err := client.Do(req)
	if err != nil {
		slog.Error("erro ao fazer a requisição", "error", err)
		// Verifica se o erro foi por timeout
		if ctx.Err() == context.DeadlineExceeded {
			slog.Error("timeout: requisição excedeu 300ms")
			return
		}
		return
	}
	defer request.Body.Close()

	// Verifica se o status code é 200 OK
	if request.StatusCode != http.StatusOK {
		slog.Error("servidor retornou erro", "status_code", request.StatusCode)
		return
	}

	// lê a requisição do server
	response, err := io.ReadAll(request.Body)
	if err != nil {
		slog.Error("erro ao ler o body da resposta", "error", err)
		return
	}

	duration := time.Since(start)
	slog.Info("requisição finalizada", "duração_ms", duration.Milliseconds())

	var bid string
	err = json.Unmarshal(response, &bid)
	if err != nil {
		slog.Error("erro ao fazer o parse da resposta", "error", err)
		return
	} else {
		slog.Info("cotação do dólar lida com sucesso", "bid", bid)
	}

	// verifica se o valor é diferenteigual a zero (vazio) antes de gravar no arquivo
	if bid == "" {
		slog.Error("cotação do dólar vazia, não será gravada no arquivo")
		return
	}

	// Cria o arquivo cotacao.txt
	file, err := os.Create("cotacao.txt")
	if err != nil {
		slog.Error("erro ao criar o arquivo", "error", err)
	}
	defer file.Close()

	if err == nil {
		_, err = file.WriteString("Dólar:{" + bid + "}")
		if err != nil {
			slog.Error("erro ao escrever no arquivo", "error", err)
		} else {
			slog.Info("cotação do dólar gravada no arquivo cotacao.txt")
		}
	}
}
