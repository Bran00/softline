# Gestão de Dados

## Descrição
Gestão de Dados é um projeto de software desenvolvido para facilitar a gestão e organização de dados em um ambiente empresarial. Este sistema permite gerenciar informações de clientes, produtos, vendas e outros aspectos relevantes para o negócio, proporcionando uma interface intuitiva e eficiente para os usuários.

## Tecnologias Utilizadas
- Frontend:
  - Next.js
  - React
  - Axios
- Backend:
  - Go
  - MySQL
  - Gorilla Mux
  - JWT Go
- Outras tecnologias:
  - Dotenv

## Configuração
## Banco de dados 

## Script do Banco de Dados

Você pode utilizar o seguinte script SQL para criar o banco de dados e as tabelas necessárias:

```sql
-- Criação do banco de dados
CREATE DATABASE IF NOT EXISTS softline;

-- Uso do banco de dados
USE softline;

-- Tabela de Clientes
CREATE TABLE IF NOT EXISTS clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    fantasia VARCHAR(255),
    documento VARCHAR(20) NOT NULL,
    endereco VARCHAR(255)
);

-- Tabela de Produtos
CREATE TABLE IF NOT EXISTS produtos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao TEXT,
    codigo_barras VARCHAR(50),
    valor_venda DECIMAL(10, 2),
    peso_bruto DECIMAL(10, 2),
    peso_liquido DECIMAL(10, 2)
);
```

### Backend
1. Clone o repositório do projeto.
2. Certifique-se de ter o Go instalado em sua máquina.
3. Instale as dependências do backend executando `go mod tidy`.
4. Configure as variáveis de ambiente no arquivo `.env` conforme necessário.
5. Inicie o servidor backend executando `go run main.go`.

### Frontend
1. Navegue até a pasta do frontend no repositório clonado.
2. Instale as dependências executando `npm install`.
3. Configure as variáveis de ambiente no arquivo `.env.local` conforme necessário.
4. Inicie o servidor de desenvolvimento executando `npm run dev`.

## Variáveis de Ambiente
Certifique-se de configurar corretamente as seguintes variáveis de ambiente antes de iniciar o servidor backend ou frontend:

- `DB_HOST`: Endereço do host do banco de dados MySQL.
- `DB_PORT`: Porta do banco de dados MySQL.
- `DB_USER`: Nome de usuário do banco de dados MySQL.
- `DB_PASSWORD`: Senha do usuário do banco de dados MySQL.
- `DB_NAME`: Nome do banco de dados MySQL.
- `PASSWORD`: Senha específica para autenticação na aplicação.
- `ADMIN_USER`: Nome de usuário do administrador da aplicação.
- `SECRET_KEY`: Chave secreta utilizada para várias finalidades na aplicação.

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues para reportar problemas ou propor novas funcionalidades. Pull requests também são encorajados.

## Licença
Este projeto está licenciado sob a [Licença MIT](https://opensource.org/licenses/MIT).

---
