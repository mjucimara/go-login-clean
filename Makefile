# =========================
# Configurações básicas
# =========================

APP_NAME := go-login-clean
CMD_DIR := ./cmd/api

# =========================
# Alvos padrão
# =========================

.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo "  make run        - Executa a aplicação"
	@echo "  make test       - Executa todos os testes"
	@echo "  make build      - Compila o binário"
	@echo "  make fmt        - Formata o código (gofmt)"
	@echo "  make vet        - Executa go vet"
	@echo "  make lint       - Verificações básicas (fmt + vet)"
	@echo "  make clean      - Remove arquivos gerados"

# =========================
# Execução
# =========================

.PHONY: run
run:
	go run $(CMD_DIR)

# =========================
# Build
# =========================

.PHONY: build
build:
	go build -o bin/$(APP_NAME) $(CMD_DIR)

# =========================
# Testes
# =========================

.PHONY: test
test:
	go test ./...

# =========================
# Qualidade de código
# =========================

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: fmt vet

# =========================
# Limpeza
# =========================

.PHONY: clean
clean:
	rm -rf bin
