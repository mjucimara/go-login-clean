package usuario

// Usuario representa uma pessoa que pode se autenticar no sistema.
type Usuario struct {
	ID        string
	Email     string
	HashSenha string
	Ativo     bool
}
