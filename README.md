---

# Movie API

Bem-vindo à **Movie API**, um projeto simples desenvolvido em **Go** utilizando o framework **Gin** e organizado com **Arquitetura Limpa (Clean Architecture)**. Esta API gerencia informações de filmes, permitindo listar, criar, atualizar, buscar e deletar filmes.

Este projeto foi criado como parte do meu portfólio para demonstrar habilidades em desenvolvimento de APIs RESTful em Go.

---

## **Endpoints**

| Método  | Rota           | Descrição                             |
|---------|----------------|---------------------------------------|
| `GET`   | `/movies`      | Lista todos os filmes.               |
| `GET`   | `/movies/:id`  | Retorna os detalhes de um filme.      |
| `POST`  | `/movies`      | Cria um novo filme.                  |
| `PUT`   | `/movies/:id`  | Atualiza as informações de um filme. |
| `DELETE`| `/movies/:id`  | Remove um filme.                     |

---

## **Estrutura do Projeto**

O projeto segue a **Arquitetura Limpa**, com separação de responsabilidades em diferentes camadas:

```
movie-api/
├── cmd/                 # Arquivo principal da aplicação (main.go)
├── internal/
│   ├── delivery/        # Camada responsável pelas rotas (HTTP via Gin)
│   ├── domain/          # Entidades e contratos (interfaces)
│   ├── repository/      # Implementação de persistência (SQL Server)
│   ├── usecase/         # Lógica de negócios (casos de uso)
├── go.mod               # Arquivo de definição de dependências
└── go.sum               # Hashes das dependências
```

---

## **Tecnologias Utilizadas**

- **Linguagem:** Go (Golang)
- **Framework:** Gin
- **Arquitetura:** Clean Architecture
- **Persistência:** SQL Server

---

## **Como Executar o Projeto**

1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/movie-api.git
   cd movie-api
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

   A conexão com o SQL Server é configurada no arquivo `cmd/main.go`. Por padrão
   são utilizados os seguintes dados:

   - **Servidor:** `dtc.erp-pegasus.com.br`
   - **Porta:** `7557`
   - **Usuário:** `rpa_bi_rwu`
   - **Senha:** `9zpzpYoi`
   - **Banco:** `RPA_BI`

3. Execute a aplicação:
   ```bash
   go run cmd/main.go
   ```

4. Acesse os endpoints:
   - Base URL: `http://localhost:8080`
   - Utilize ferramentas como **Postman**, **cURL** ou um navegador para testar.

---

## **Exemplos de Uso**

### **Criar um Filme**
- **Rota:** `POST /movies`
- **Body (JSON):**
  ```json
  {
      "title": "Inception",
      "director": "Christopher Nolan",
      "genre": "Sci-Fi",
      "year": 2010
  }
  ```
- **Resposta (201):**
  ```json
  {
      "id": 1,
      "title": "Inception",
      "director": "Christopher Nolan",
      "genre": "Sci-Fi",
      "year": 2010
  }
  ```

### **Listar Todos os Filmes**
- **Rota:** `GET /movies`
- **Resposta (200):**
  ```json
  [
      {
          "id": 1,
          "title": "Inception",
          "director": "Christopher Nolan",
          "genre": "Sci-Fi",
          "year": 2010
      }
  ]
  ```

### **Buscar um Filme por ID**
- **Rota:** `GET /movies/1`
- **Resposta (200):**
  ```json
  {
      "id": 1,
      "title": "Inception",
      "director": "Christopher Nolan",
      "genre": "Sci-Fi",
      "year": 2010
  }
  ```

### **Atualizar um Filme**
- **Rota:** `PUT /movies/1`
- **Body (JSON):**
  ```json
  {
      "title": "Inception Updated",
      "director": "Christopher Nolan",
      "genre": "Sci-Fi",
      "year": 2010
  }
  ```
- **Resposta (200):**
  ```json
  {
      "id": 1,
      "title": "Inception Updated",
      "director": "Christopher Nolan",
      "genre": "Sci-Fi",
      "year": 2010
  }
  ```

### **Deletar um Filme**
- **Rota:** `DELETE /movies/1`
- **Resposta (204):**
  - Sem conteúdo.

---
