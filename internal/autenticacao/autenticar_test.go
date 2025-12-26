package autenticacao

import (
	"testing"

	"github.com/mjucimara/go-login-clean/internal/armazenamento"
	"github.com/mjucimara/go-login-clean/internal/usuario"
)

func TestAutenticar_ComCredenciaisInvalidas_DeveRetornarErro(t *testing.T) {
	memoria := armazenamento.NovaMemoria(map[string]usuario.Usuario{})
	repositorio := usuario.NovoRepositorio(memoria)
	autenticador := NovoAutenticador(repositorio)

	_, err := autenticador.Autenticar("invalido@exemplo.com", "errada")
	if err == nil {
		t.Fatal("esperava erro para credenciais inválidas")
	}
}
