# go-login-clean

Login web simples em Go, com frontend estático servido via `net/http`, arquitetura limpa e foco em clareza, previsibilidade e boa experiência de uso.

Este projeto **não busca complexidade**, mas sim demonstrar **boas decisões técnicas**, separação de responsabilidades e integração correta entre frontend e backend usando apenas a biblioteca padrão do Go.

---

## 🎯 Objetivo do projeto

Demonstrar:

* Como servir arquivos estáticos corretamente em Go
* Como estruturar um backend simples e limpo
* Como integrar um frontend existente sem reescrevê-lo
* Como tratar HTTP de forma correta (métodos, status, JSON)
* Como cuidar de UX sem exageros ou frameworks

---

## 🧠 Decisões técnicas

### Por que Go stdlib?

* Menos dependências
* Comportamento previsível
* Código mais fácil de entender e manter
* Ideal para sistemas pequenos e confiáveis

### Por que frontend estático?

* Não exige build
* Pode ser servido diretamente pelo backend
* Facilita manutenção e deploy
* Evita complexidade desnecessária

### Por que não há banco de dados?

* O foco é arquitetura e fluxo, não persistência
* Usuários são armazenados em memória
* Facilita testes e entendimento do domínio

### Por que não há JWT ou autenticação real?

* Não é o objetivo do projeto
* Evita desviar atenção para detalhes de segurança
* Mantém o escopo controlado

---

## 📁 Estrutura do projeto

```
go-login-clean/
├── cmd/
│   └── api/              # Ponto de entrada da aplicação
│       └── main.go
├── internal/
│   ├── autenticacao/     # Lógica de autenticação
│   ├── usuario/          # Domínio de usuário
│   ├── armazenamento/   # Armazenamento em memória
│   └── httpapi/          # Handlers HTTP
├── web/
│   ├── index.html        # Frontend
│   ├── css/
│   │   └── styles.css
│   └── js/
│       └── index.js
├── Makefile
├── go.mod
└── README.md
```

---

## ▶️ Como executar

### Pré-requisitos

* Go 1.21+ (ou compatível)

### Executar a aplicação

```bash
make run
```

Ou, sem Makefile:

```bash
go run ./cmd/api
```

Acesse no navegador:

```
http://localhost:8080
```

---

## 🧪 Testes

Executar todos os testes:

```bash
make test
```

Os testes focam na **lógica de domínio**, sem dependência de HTTP ou frontend.

---

## 🔌 API

### POST `/login`

**Request**

```json
{
  "email": "user@exemplo.com",
  "senha": "segredo"
}
```

**Response (sucesso)**

```json
{
  "token": "token-falso"
}
```

**Response (erro)**

```json
{
  "erro": "email_ou_senha_invalidos"
}
```

Status HTTP apropriados são utilizados (`200`, `401`, `405`).

---

## 🎨 Frontend

* Interface simples e limpa
* Servida diretamente pelo backend
* Sem frameworks JS
* UX com:

  * foco automático
  * estado de carregamento
  * mensagens claras
  * dark mode automático (via CSS)

---

## 🚫 O que este projeto **não** tenta ser

* Um sistema de autenticação completo
* Um exemplo de segurança em produção
* Um template frontend moderno
* Um projeto com banco de dados
* Um SaaS ou aplicação escalável

Essas decisões são **intencionais**.

---

## 🧠 Filosofia

> Código claro é mais valioso que código impressionante.
> Um projeto pequeno e bem finalizado ensina mais do que um grande e inacabado.

---

## 📌 Próximos passos possíveis

* Persistência real (banco de dados)
* Autenticação com tokens reais
* Proteção de rotas
* CLI complementar
* Projeto separado focado em API REST

---

## 📄 Licença

Este projeto é apenas demonstrativo e educacional.

---

Se quiser, posso:

* revisar o README como se fosse code review
* encurtar para versão “portfólio”
* adaptar para clientes ou recrutadores
* encerrar oficialmente o projeto

É só dizer.
