package autenticacao

import (
	"errors"

	"github.com/mjucimara/go-login-clean/internal/usuario"
)

var ErroCredenciaisInvalidas = errors.New("credenciais inválidas")

type Autenticador struct {
	usuarios *usuario.Repositorio
}

func NovoAutenticador(r *usuario.Repositorio) *Autenticador {
	return &Autenticador{usuarios: r}
}

func (a *Autenticador) Autenticar(email, senha string) (string, error) {
	u, encontrado := a.usuarios.BuscarPorEmail(email)
	if !encontrado || !u.Ativo {
		return "", ErroCredenciaisInvalidas
	}

	if !senhaCorresponde(senha, u.HashSenha) {
		return "", ErroCredenciaisInvalidas
	}

	return "token-falso", nil
}

func senhaCorresponde(senha, hash string) bool {
	return senha == "segredo"
}
