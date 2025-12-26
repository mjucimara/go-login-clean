package usuario

// Buscador define o comportamento mínimo para buscar usuários.
type Buscador interface {
	BuscarUsuarioPorEmail(email string) (Usuario, bool)
}

// Repositorio encapsula o acesso aos usuários.
type Repositorio struct {
	buscador Buscador
}

func NovoRepositorio(b Buscador) *Repositorio {
	return &Repositorio{buscador: b}
}

func (r *Repositorio) BuscarPorEmail(email string) (Usuario, bool) {
	return r.buscador.BuscarUsuarioPorEmail(email)
}
