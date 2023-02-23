package main

import (
	"fmt"
	"modulo/helper"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	helper.Write()
	erro := checkmail.ValidateFormat("nollhyury@gmail.com")
	fmt.Println(erro)
}
