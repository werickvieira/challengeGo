package main

import (
	"fmt"
	"strings"
)

// Coordena model
type Coordena struct {
	X int
	Y int
}

var pontos = []Coordena{{ 0, 0 }}
var coordenadaY = 0
var coordenadaX = 0

func controlaCoordenada(arr []Coordena, novo Coordena) bool {
	fmt.Println("@@@ CONTROLA COORDENADA @@@")
	fmt.Println("@@@ COORDENADA NOVA", novo)
	fmt.Println("@@@ LISTA NOVA", arr)
	fmt.Println("")
	for _, elemento := range arr {
		if elemento.X == novo.X && elemento.Y == novo.Y {
			fmt.Println("PONTO UTILIZADO!")
			return true
		}
	}
	return false
}

func comandos() string {
	var input string
	fmt.Println("C - [cima] | D - [direita] | B - [baixo] | E - [esquerda]")
	fmt.Println("Entre com o comando: ")
	fmt.Scanln(&input)
	char := strings.ToUpper(input)
	return char
}

func insereCoordenada(status bool, char string) {
	if (!status) {
		if (char == "C") { coordenadaY++ }
		if (char == "B") { coordenadaY-- }
		if (char == "D") { coordenadaX++ }
		if (char == "E") { coordenadaX-- }
		novaCoordena := Coordena{coordenadaY, coordenadaX}
		fmt.Println("@@@ INSERIR COORDENADA @@@")
		fmt.Println("@@@ INSERIR", novaCoordena)
		fmt.Println("@@@ PONTOS", pontos)
		fmt.Println("")
		permitido := controlaCoordenada(pontos, novaCoordena)
		fmt.Println("@@@ PERMITIDO", permitido)
		fmt.Println("")
		pontos = append(pontos, novaCoordena)
		if (!permitido) {
			// LOOP
			localchar := comandos()
			insereCoordenada(permitido, localchar)
		}
	}
}

func main() {
	char := comandos()
	insereCoordenada(false, char)
	fmt.Println("@@@ PONTOS ATUALIZADOS", pontos)
}