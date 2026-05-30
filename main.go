package main

import (
	"time"

	"github.com/denispolastri/desafio-client-server-api/client"
	"github.com/denispolastri/desafio-client-server-api/server"
)

func main() {
	// Inicia o servidor em uma thread separada para não bloquear a execução do cliente
	go server.Server()

	// Aguarda o servidor iniciar
	time.Sleep(3 * time.Second)

	// Inicia o cliente
	client.Client()
}
