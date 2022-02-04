package main

import (
	"fmt"
	emails "testes-automatizados/emailsValidos"
)

func main() {
	tipoEmail := emails.TiposDeEmails("victor@gmail.com")
	fmt.Println(tipoEmail)
}
