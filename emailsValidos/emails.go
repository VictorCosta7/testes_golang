package emails

import (
	"strings"
)

//Função para representar testes em multiplos cenários
//TiposDeEmails verifica se o email inserido é válido
func TiposDeEmails(email string) string {
	emailsValidos := []string{"gmail.com", "gmail.com.br", "hotmail.com", "hotmail.com.br", "yahoo.com", "yahoo.com.br"}
	emailEmMinusculo := strings.ToLower(email)
	coletarPalavraValida := strings.Split(emailEmMinusculo, "@")[1]

	emailTemUmTipoValido := false
	for _, tipo := range emailsValidos {
		if tipo == coletarPalavraValida {
			emailTemUmTipoValido = true
		}
	}

	if emailTemUmTipoValido {
		return coletarPalavraValida
	}
	return "Email inválido"
}
