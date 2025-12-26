package armazenamento

import "github.com/mjucimara/go-login-clean/internal/usuario"

// Memoria armazena usuários em memória.
// Usado apenas para desenvolvimento e testes iniciais.
type Memoria struct {
	usuarios map[string]usuario.Usuario
}

func NovaMemoria(usuarios map[string]usuario.Usuario) *Memoria {
	return &Memoria{usuarios: usuarios}
}

// BuscarUsuarioPorEmail retorna um usuário pelo email.
func (m *Memoria) BuscarUsuarioPorEmail(email string) (usuario.Usuario, bool) {
	u, encontrado := m.usuarios[email]
	return u, encontrado
}
