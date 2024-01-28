# Projeto Ikaros

## Sobre o Projeto

O Projeto Ikaros é uma aplicação back-end desenvolvida em Go, integrada com MongoDB, destinada a demonstrar práticas avançadas em programação e design de API. Este projeto enfoca eficiência, segurança e escalabilidade, oferecendo uma solução robusta para gerenciamento de dados.

## Características

- **Linguagem Go**: Utiliza Go para desempenho e manutenção eficientes.
- **API RESTful**: Construída com o framework Gorilla Mux.
- **Integração com MongoDB**: Para operações CRUD em dados cadastrais.
- **Autenticação e Segurança**: Com base em tokens e práticas robustas de segurança.
- **Prevenção de Duplicidade**: Garantia de unicidade de dados.
- **Arquivamento de Dados**: Estratégia segura de "deleção" de dados.

## Funcionalidades

- **CRUD Completo**: Leitura, inserção, atualização e deleção (arquivamento) de dados.
- **Verificação de Duplicidade**: Para evitar dados repetidos.
- **Segurança**: Proteção robusta dos dados e das operações da API.

## Pré-Requisitos

- Go (versão mais recente)
- MongoDB
- Git (opcional, para clonagem do repositório)

## Instalação e Configuração

1. **Clone o Repositório (opcional)**:
   ```
   git clone https://seu-repositorio/ikaros.git
   cd ikaros
   ```
2. **Instale as Dependências**:
   ```
   go mod tidy
   ```
3. **Inicie o MongoDB**:
   - Certifique-se de que o MongoDB está rodando localmente ou ajuste a `MONGODB_URI`.

## Execução

Execute o servidor com:

```
go run backend/api/main.go
```

O servidor estará disponível na porta configurada (padrão: 8080).

## Testando a API

Utilize ferramentas como Postman ou `curl` para testar as funcionalidades da API.

## Roadmap

### Concluído

- Fundamentos da Linguagem Go
- Leitura e Escrita de Arquivos
- Lógica de Negócios Inicial
- API REST com Gorilla Mux
- Autenticação e Segurança

### Em Andamento

- Integração com MongoDB

### Próximos Passos

- Desenvolvimento Frontend
- Scripting e Análise de Dados
- Deploy e Manutenção
- Documentação Completa
- Projetos Complementares

## Como Contribuir

Contribuições são bem-vindas. Para contribuir:

- Faça um fork do projeto.
- Crie uma branch para suas alterações.
- Envie um Pull Request.

Para reportar bugs ou sugerir melhorias, abra uma issue.
