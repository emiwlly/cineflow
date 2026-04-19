# Cineflow

Um gerenciador de filmes desenvolvido em Go com interface de linha de comando (CLI). Permite gerenciar um catálogo de filmes com operações CRUD completas.

## 📋 Descrição

Cineflow é uma aplicação simples e intuitiva para gerenciar um acervo de filmes. Com ela você pode:

- **Listar** todos os filmes do catálogo
- **Adicionar** novos filmes
- **Atualizar** informações dos filmes
- **Remover** filmes do catálogo

Cada filme contém:
- **ID**: Identificador único
- **Nome**: Título do filme
- **Quantidade**: Número de cópias disponíveis

## 🛠️ Tecnologias

- **Linguagem**: Go 1.21
- **Containerização**: Docker
- **Orquestração**: Docker Compose
- **Interface**: CLI (Linha de Comando)

## 📁 Estrutura do Projeto

```
cineflow/
├── main/
│   └── main.go           # Arquivo principal da aplicação
├── models/               # Estruturas de dados (em desenvolvimento)
├── services/             # Serviços da aplicação (em desenvolvimento)
├── Dockerfile            # Configuração para containerização
├── docker-compose.yml    # Orquestração de containers
└── README.md            # Este arquivo
```

## 📝 Funcionalidades

- Gerenciamento de catálogo de filmes
- Operações CRUD completas
- Interface interativa via CLI
- Execução containerizada com Docker

## 🔄 Operações Disponíveis

A aplicação oferece as seguintes operações:
1. **Visualizar filmes** - Exibe todos os filmes cadastrados
2. **Adicionar filme** - Cria novo entrada no catálogo
3. **Atualizar filme** - Modifica informações de um filme existente
4. **Deletar filme** - Remove um filme do catálogo

## 📦 Arquitetura

O projeto segue uma arquitetura modular com:
- `main/`: Lógica principal e orquestração
- `models/`: Estruturas de dados (estrutura preparada)
- `services/`: Serviços e utilitários (estrutura preparada)

---

**Autor**: Estudo de Go  
**Status**: Em desenvolvimento