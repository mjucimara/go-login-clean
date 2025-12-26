const form = document.getElementById("login-form")
const botao = document.getElementById("btn-login")
const mensagem = document.getElementById("mensagem")

form.addEventListener("submit", function (e) {
  e.preventDefault()

  mensagem.textContent = ""

  const email = document.getElementById("email").value
  const senha = document.getElementById("senha").value

  const textoOriginal = botao.textContent
  botao.disabled = true
  botao.textContent = "Entrando..."

  fetch("/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, senha })
  })
    .then(async res => {
      const dados = await res.json()

      if (!res.ok) {
        throw new Error(dados.erro || "erro_desconhecido")
      }

      return dados
    })
    .then(() => {
      mensagem.style.color = "#1c7ed6"
      mensagem.textContent = "Login realizado com sucesso"
    })
    .catch(err => {
      mensagem.style.color = "#e03131"

      if (err.message === "email_ou_senha_invalidos") {
        mensagem.textContent = "Email ou senha incorretos"
      } else {
        mensagem.textContent = "Erro inesperado. Tente novamente."
      }
    })
    .finally(() => {
      botao.disabled = false
      botao.textContent = textoOriginal
    })
})

/* === Mostrar / esconder senha === */

const senhaInput = document.getElementById("senha")
const eye = document.getElementById("eye")
const eyeSlash = document.getElementById("eye-slash")

eyeSlash.addEventListener("click", () => {
  senhaInput.type = "text"
  eyeSlash.style.display = "none"
  eye.style.display = "block"
})

eye.addEventListener("click", () => {
  senhaInput.type = "password"
  eye.style.display = "none"
  eyeSlash.style.display = "block"
})

/* === Foco automático no email === */

window.addEventListener("load", () => {
  document.getElementById("email").focus()
})
