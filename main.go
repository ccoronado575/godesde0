package main

import (
	"fmt"

	"github.com/ccoronado575/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1223)
	fmt.Println(estado)
	fmt.Println(texto)

}
