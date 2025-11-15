# Plataforma de Upskilling/Reskilling API

|Nome|RM|
|:-:|:-:|
|Matheus Zottis|94119|
|Victor Didoff|552965|
|Vinicius Silva|553240|

## 🛠️ Tecnologias Utilizadas

*   **Linguagem:** Go (Golang) - Versão 1.24+
*   **Framework Web:** Gin-Gonic
*   **Banco de Dados:** PostgreSQL
*   **Driver DB:** `github.com/lib/pq`
*   **Orquestração:** Docker e Docker Compose

## 🚀 Como Executar o Projeto

O projeto é configurado para ser executado facilmente via Docker Compose, que gerencia o banco de dados PostgreSQL e a aplicação Go.

### Pré-requisitos

*   Docker
*   Docker Compose

### 1. Configuração

Crie um arquivo `.env` na raiz do projeto, copiando o conteúdo de `.env.example`.

```bash
cp .env.example .env
```

O conteúdo padrão do `.env` será:

```
# Configurações do Servidor
PORT=8080

# Configurações do Banco de Dados PostgreSQL
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=upskilling_db
```

### 2. Execução

Execute o comando abaixo para construir as imagens, iniciar os containers e rodar o seeder (setup) automaticamente:

```bash
docker compose up -d
```

A aplicação estará acessível em `http://localhost:8080`.

### 3. Migrações e Seeds

As migrações (criação das tabelas) e o seeder (população inicial de dados) são executados automaticamente pelo container da aplicação (`app`) antes de iniciar o servidor.

O script de setup (`cmd/start/setup.go`) garante que:
1.  A conexão com o PostgreSQL seja estabelecida.
2.  As tabelas sejam criadas (baseado em `db/schema.sql`).
3.  Dados iniciais de `Usuários`, `Trilhas`, `Competências` e `Matrículas` sejam inseridos, caso as tabelas estejam vazias.

## 🔗 Endpoints da API (v1)

A API expõe os seguintes endpoints sob o prefixo `/api/v1`:

| Recurso | Método | URL | Descrição |
| :--- | :--- | :--- | :--- |
| **Usuários** | `POST` | `/api/v1/usuarios` | Cria um novo usuário. |
| | `GET` | `/api/v1/usuarios` | Lista todos os usuários. |
| | `GET` | `/api/v1/usuarios/{id}` | Busca usuário por ID. |
| | `PUT` | `/api/v1/usuarios/{id}` | Atualiza usuário por ID. |
| | `DELETE` | `/api/v1/usuarios/{id}` | Deleta usuário por ID. |
| **Trilhas** | `POST` | `/api/v1/trilhas` | Cria uma nova trilha. |
| | `GET` | `/api/v1/trilhas` | Lista todas as trilhas. |
| | `GET` | `/api/v1/trilhas/{id}` | Busca trilha por ID. |
| | `PUT` | `/api/v1/trilhas/{id}` | Atualiza trilha por ID. |
| | `DELETE` | `/api/v1/trilhas/{id}` | Deleta trilha por ID. |
| **Matrículas** | `POST` | `/api/v1/matriculas` | Matricular usuário em uma trilha. |
| | `GET` | `/api/v1/usuarios/{id}/matriculas` | Lista matrículas de um usuário. |

### Exemplo de Requisição (Criação de Usuário)

**URL:** `POST http://localhost:8080/api/v1/usuarios`

**Body (JSON):**

```json
{
  "nome": "João da Silva",
  "email": "joao.silva@exemplo.com",
  "area_atuacao": "Engenharia de Software",
  "nivel_carreira": "Pleno"
}
```

**Resposta (201 Created):**

```json
{
  "id": 5,
  "nome": "João da Silva",
  "email": "joao.silva@exemplo.com",
  "area_atuacao": "Engenharia de Software",
  "nivel_carreira": "Pleno",
  "data_cadastro": "2025-11-12T10:30:00Z"
}
```

## 📝 Documentação (Swagger)

esse passo a passo ja foi feito, so precisa ser realizado se desejar atualizar algum dado.

A documentação Swagger foi configurada no código-fonte (comentários `// @...`) para ser gerada localmente.

Para gerar a documentação na sua máquina, instale o `swag`:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

E execute na raiz do projeto:

```bash
swag init -g upskilling-server.go
```

Isso criará a pasta `docs` e os arquivos necessários. A rota para a documentação será: `http://localhost:8080/swagger/index.html` (após descomentar a linha no `upskilling-server.go`).
