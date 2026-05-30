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

Existem dois pontos de testes no server.go, na linha 131, descomentando o sleep(), irá simular a demora de gravação do dolar no banco de dados, o que vai gerar um erro de timeout.

Na linha 161, outro ponto de sleep(), para simular a demora de resposta do servidor.

Em ambos os casos, o arquivo cotacao.txt não é criado no client.go.
