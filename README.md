# desafio-client-server-api
Desafio para aplicar conhecimentos de HTTP, Contextos, Banco de Dados e Manipulação de Arquivos em Go.

### Instruções do projeto
```
|_client
|   client.go
|_server
|   server.go
|.gitignore
| dolar.db
| go.mod
| go.sum
| main.go
```

O arquivo main.go deve ser executado na raiz do projeto, na pasta desafio-client-server-api. Após a execução do server.go, é executado um sleep de 3 segundos para a chamada do client.go.

O server.go cria o banco de dados, a tabela dolar através da func initDb() e aguarda chamadas HTTP através do endereço http://localhost:8080/cotacao

O client faz um Get no endereço localhost:8080/cotacao e chama a função consultaCotacaoSiteEconomia() que consulta o valor do dolar e grava no banco de dados através da func gravaCotacao().

O client recebe o valor do dolar no request.Body e grava o valor no arquivo cotacao.txt através da func gravaArqTxt().