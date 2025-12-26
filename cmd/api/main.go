package main

import (
	"log"
	"net/http"

	"github.com/mjucimara/go-login-clean/internal/armazenamento"
	"github.com/mjucimara/go-login-clean/internal/autenticacao"
	"github.com/mjucimara/go-login-clean/internal/httpapi"
	"github.com/mjucimara/go-login-clean/internal/usuario"
)

func main() {
	memoria := armazenamento.NovaMemoria(map[string]usuario.Usuario{
		"user@exemplo.com": {
			ID:        "1",
			Email:     "user@exemplo.com",
			HashSenha: "hash",
			Ativo:     true,
		},
	})

	repositorio := usuario.NovoRepositorio(memoria)
	autenticador := autenticacao.NovoAutenticador(repositorio)
	loginHandler := httpapi.NovoLoginHandler(autenticador)

	mux := http.NewServeMux()

	// ⚠️ FRONTEND
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	// ⚠️ API
	mux.Handle("/login", loginHandler)

	log.Println("servidor iniciado em :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
