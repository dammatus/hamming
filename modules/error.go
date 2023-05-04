package modules

import (
	"math/rand"
	"time"
)

func GenerarErrorEnbloque(encoded []byte) []byte {

	// Semilla del generador de numeros aleatorios
	rand.Seed(time.Now().UnixNano())

	// Generar un número aleatorio entre 0 y blockSize -> [0,blockSize)
	num := rand.Intn(len(encoded))

	// Si en num hay un 0 -> lo cambia a 1. Si hay un 1 -> Lo cambia a 0
	encoded[num] = encoded[num] ^ 1

	return encoded
}
