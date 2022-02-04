package emails

import "testing"

type cenariosDeTestes struct {
	emailParaTeste    string
	resultadoEsperado string
}

func TestTiposDeEmails(t *testing.T) {

	cenariosDeTestes := []cenariosDeTestes{
		{"victor@hotmail.com", "hotmail.com"},
		{"victorcosta10@gmail.com", "gmail.com"},
		{"victor@hotmail.com.br", "hotmail.com.br"},
		{"victor@gmail.com", "gmail.com"},
		{"victorr@gmail.com", "gmail.com"},
		{"victor@yahoo.com", "yahoo.com"},
		{"victor@yahoo.com.br", "yahoo.com.br"},
		{"victor@gmailll.com", "Email inválido"},
		{"victor@Gmail.com", "gmail.com"},

		//Posso acrescetar diversos cenários
	}

	for _, cenarios := range cenariosDeTestes {
		tipoDeEmailRecebido := TiposDeEmails(cenarios.emailParaTeste)
		if tipoDeEmailRecebido != cenarios.resultadoEsperado {
			t.Errorf("O email recebido %s é diferente do email esperado %s", tipoDeEmailRecebido, cenarios.resultadoEsperado)
		}
	}
}
