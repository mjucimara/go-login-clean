package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/mjucimara/go-login-clean/internal/autenticacao"
)

type loginHandler struct {
	autenticador *autenticacao.Autenticador
}

func NovoLoginHandler(a *autenticacao.Autenticador) http.Handler {
	return &loginHandler{autenticador: a}
}

type requisicaoLogin struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type respostaLogin struct {
	Token string `json:"token"`
}

type respostaErro struct {
	Erro string `json:"erro"`
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req requisicaoLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		escreverErro(w, http.StatusBadRequest, "json_invalido")
		return
	}

	token, err := h.autenticador.Autenticar(req.Email, req.Senha)
	if err != nil {
		escreverErro(w, http.StatusUnauthorized, "email_ou_senha_invalidos")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respostaLogin{
		Token: token,
	})
}

func escreverErro(w http.ResponseWriter, status int, mensagem string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(respostaErro{
		Erro: mensagem,
	})
}
