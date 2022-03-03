package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int16
}

func main() {
	fmt.Println("Arquivo Struct")

	var u usuario
	u.nome = "Cláudio"
	u.idade = 47

	enderecoExemplo := endereco{"Rua José de Alencar", 436}

	fmt.Println(u)
	fmt.Println(u.nome)
	fmt.Println(u.idade)

	u2 := usuario{"Aldo", 2, enderecoExemplo}
	fmt.Println(u2.nome)
	fmt.Println(u2.idade)
	fmt.Println(u2.endereco)

	u3 := usuario{nome: "Antonio"}
	fmt.Println(u3.nome)

}
