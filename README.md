# janeiro_Desafio1_Client-Server
Desafio de Janeiro 2025 Faculdade Full Cycle - Go Expert - Client_Server

### 📖 Descrição
Janeiro Desafio 1: Client-Server é uma aplicação desenvolvida em Go que consiste em um servidor e um cliente. O servidor expõe um endpoint para obter a cotação do dólar em relação ao real (USD-BRL) a partir de uma API externa, salva essa cotação em um banco de dados SQLite e também em um arquivo de texto. O cliente faz uma requisição ao servidor para obter a cotação e armazena o resultado em um arquivo local.

### 🌟 Funcionalidades
Servidor:

Exponha um endpoint /cotacao para obter a cotação do dólar.
Chama uma API externa para buscar a cotação atual do USD-BRL.
Salva a cotação em um banco de dados SQLite com timeout de 10ms.
Retorna a cotação em formato JSON para o cliente.
Gerenciamento de timeouts para chamadas à API externa (200ms) e operações de banco de dados (10ms).
Cliente:

Faz uma requisição ao servidor para obter a cotação do dólar.
Salva a cotação em um arquivo cotacao.txt.
Gerenciamento de timeout de 300ms para a requisição ao servidor.

###  🔧 Tecnologias Utilizadas
Go (Golang): Linguagem de programação principal.
SQLite: Banco de dados leve para persistência das cotações.
Context: Gerenciamento de timeouts e cancelamentos nas requisições.

###  📁 Estrutura do Projeto:



janeiro_Desafio1_Client-Server

├── client

│       └── main.go

├── server

│       └── main.go

├── cotacao.txt

├── cotacoes.db

├── go.mod

└── go.sum



##  🛠️ Pré-requisitos
Antes de começar, certifique-se de ter o seguinte instalado no seu sistema:

Go (versão 1.16 ou superior)
Git (opcional, para clonar o repositório)
SQLite (opcional, para visualizar o banco de dados; instruções abaixo)

## Instalação
###  1. Clone o Repositório

git clone https://github.com/willianfariabatista/janeiro_Desafio1_Client-Server.git

cd janeiro_Desafio1_Client-Server

###  2. Inicialize o Módulo Go e Instale as Dependências
O projeto já deve ter um arquivo go.mod. Se não tiver, inicialize o módulo:

go mod init janeiro_Desafio1_Client-Server

Em seguida, instale as dependências necessárias:

go get

####  Nota: Este projeto utiliza o driver SQLite sem CGO (modernc.org/sqlite). Certifique-se de que todas as dependências estão corretamente instaladas.

🏃‍♂️ Como Executar
### 1. Inicie o Servidor
Abra um terminal e navegue até a pasta do servidor:

Execute os comandos:
cd janeiro_Desafio1_Client-Server
go run server/main.go
Você deverá ver a seguinte mensagem indicando que o servidor está rodando:

Resposta do Servidor:
2025/01/29 20:17:01 Servidor iniciado na porta :8080


### 2. Execute o Cliente em uma seguindo terminal:

cd cd janeiro_Desafio1_Client-Server
go run client/main.go

Se tudo estiver funcionando corretamente, você verá uma mensagem semelhante a esta:

2025/01/29 21:13:14 Cotação salva em cotacao.txt com sucesso!

### 3. Verifique os Resultados

No Arquivo cotacao.txt: Deverá conter a cotação do dólar, por exemplo:

Banco de Dados cotacoes.db: Armazena todas as cotações recebidas. Para visualizar os dados:

Abra o SQLite via terminal (veja a seção Como Abrir o SQLite abaixo).

### Navegue até o diretório do projeto:

PS cd ..\janeiro_Desafio1_Client-Server

### Abra o banco de dados, executar o comando abaixo no powershell

PS ..\janeiro_Desafio1_Client-Server> sqlite3 cotacoes.db

### Dentro do SQLite CLI, liste as tabelas:

sqlite> .tables

### Consulte as cotações salvas:

sqlite> SELECT * FROM cotacoes;

### Para fechar o banco de dados:

sqlite> .exit


# 📚 Como Abrir o SQLite
Para visualizar o conteúdo do banco de dados cotacoes.db, siga estes passos:

## 1. Baixe e Instale o SQLite
Acesse o site oficial do SQLite: SQLite Download Page.

Na seção Precompiled Binaries for Windows, baixe o arquivo:

sqlite-tools-win-x64-<versão>.zip (por exemplo, sqlite-tools-win-x64-3480000.zip).
Extraia o conteúdo do ZIP para uma pasta fixa, como:

C:\sqlite3\sqlite-tools-win-x64-3480000

## 2. Adicione o SQLite ao PATH do Windows
Para poder usar o comando sqlite3 de qualquer lugar no terminal:

Pressione Win + R, digite sysdm.cpl e pressione Enter.

Na aba Avançado, clique em Variáveis de Ambiente.

Em Variáveis do Sistema, encontre a variável Path e clique em Editar.

Clique em Novo e adicione o caminho onde você extraiu o SQLite, por exemplo:

C:\sqlite3\sqlite-tools-win-x64-3480000
Clique em OK para salvar e feche todas as janelas.

## 3. Teste a Instalação
 Abra um novo terminal e verifique a versão do SQLite:

sqlite3 --version

### Deverá exibir algo como:

3.48.0 2024-01-16 12:34:56

## 4. Acesse o Banco de Dados
 Navegue até a pasta do projeto:

cd ..\janeiro_Desafio1_Client-Server

### Abra o banco de dados:

..\janeiro_Desafio1_Client-Server> sqlite3 cotacoes.db

### Dentro do SQLite CLI, liste as tabelas:

sqlite> .tables

### Consulte as cotações salvas:

sqlite> SELECT * FROM cotacoes;

### Para sair do SQLite:

sqlite> .exit


## 📝 Considerações Finais
Timeouts:

API de Cotação: 200ms.
Persistência no Banco: 10ms.
Cliente: 300ms.
Driver SQLite: Utilizado modernc.org/sqlite para evitar dependências com CGO.

Logs: Tanto o servidor quanto o cliente registram operações importantes e erros para facilitar o debug e o monitoramento.

## 📈 Melhorias Futuras
Endpoint de Histórico: Adicionar um endpoint para listar todas as cotações salvas no banco.
Interface Web: Criar uma interface web para visualizar as cotações em tempo real.
Teste Automatizado: Implementar testes automatizados para garantir a robustez do sistema.
Documentação Adicional: Melhorar a documentação com diagramas de arquitetura e fluxo de dados.

### 🛠️ Como Contribuir

Fork este repositório.

### Crie uma branch para sua feature:

git checkout -b minha-feature

### Commit suas mudanças:

git commit -m "Adicionar nova feature"

### Push para a branch:

git push origin minha-feature

### Abra um Pull Request.

## Contato
Se tiver dúvidas ou sugestões, sinta-se à vontade para abrir uma Issue ou entrar em contato comigo.
