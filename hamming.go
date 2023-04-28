package hamming

import (
	"fmt"
	"io/ioutil"
)

const (
	bitsParity32    = 5
	bitsParity2048  = 11
	bitsParity65536 = 16
	bitsInfo32      = 26
	bitsInfo2048    = 2036
	bitsInfo65536   = 65519
)

// Convierte de binario a texto
func binToASCII(bin []byte) string {
	/*
	   devuelve los mismo que la funcion de abajo, solo que de tipo string
	*/
	ascii := ""
	for i := 0; i < len(bin); i += 8 { // se recorre el slice de 8 bits en 8
		end := i + 8
		if end > len(bin) {
			end = len(bin)
		}
		bits := bin[i:end]
		n := byte(0)                     // inicializo un byte en cero que sera usado para construir el byte ASCII correspondiente
		for j := 0; j < len(bits); j++ { // para cada uno, se desplaza n un bit hacia la izquierda y se verifica si el bit es 1 o 0
			n <<= 1
			if bits[j] == 0x01 { // Si el bit es 1, se utiliza el operador "|" para poner el último bit de "n" en 1.
				n |= 1
			}
		}
		ascii += string(n) // se agrega el byte n al string ascii
	}
	return ascii
}

// Convierte de bytes a bits
func byteToBits(slice []byte, blockSize int) []byte {
	/*
		Convierte un slice de bytes en una secuencia de bits.
		Para cada byte del slice, toma sus bits de izquierda a derecha
		y los agrega a un slice de bytes llamado "bits"
	*/

	bits := make([]byte, 0, len(slice)*8)

	for _, b := range slice {
		for i := 7; i >= 0; i-- {
			bit := (b >> uint(i)) & 1
			bits = append(bits, bit)
		}
	}
	return bits
}

// Abre el archivo a codificar
func readFile(file string) string {
	// Lee el contenido del archivo
	datos, err := ioutil.ReadFile(file)
	if err != nil {
		// Si ocurre un error, devuelve una cadena vacía
		return ""
	}
	// Convierte el slice de bytes a un string y lo devuelve
	return string(datos)
}

func bitsToByte(bits []byte) []byte {
	/*
		Convierte una secuencia de bits en un slice de bytes.
		Agrupa los bits de 8 en 8 y los convierte en un byte.
	*/

	// Aseguramos que la longitud de bits sea múltiplo de 8
	numBits := len(bits)
	if numBits%8 != 0 {
		pad := make([]byte, 8-(numBits%8))
		bits = append(bits, pad...)
	}

	// Convertimos los bits en bytes
	numBytes := len(bits) / 8
	bytes := make([]byte, numBytes)
	for i := 0; i < numBytes; i++ {
		for j := 0; j < 8; j++ {
			bit := bits[(i*8)+j]
			bytes[i] = (bytes[i] << 1) | bit
		}
	}
	return bytes
}

// Escribe en un archivo la codificacion
func writeFile(file string, datos string) error {
	// Escribe el contenido en el archivo
	err := ioutil.WriteFile(file, []byte(datos), 0644)
	if err != nil {
		// Si ocurre un error, devuelve el error
		return err
	}
	// Si no hay errores, devuelve nil
	return nil
}
func main() {

	var blockSize int
	fmt.Println("Ingrese un numero entero:")
	_, err := fmt.Scan(&blockSize)
	if err != nil {
		fmt.Println("Error al leer el numero:", err)
		return
	}
	var parityBits int
	var infoBits int
	switch blockSize {
	case 32:
		parityBits = bitsParity32
		infoBits = bitsInfo32
	case 2048:
		parityBits = bitsParity2048
		infoBits = bitsInfo2048
	case 65536:
		parityBits = bitsParity65536
		infoBits = bitsInfo65536
	default:
		fmt.Print("Error")
	}
	// STRING
	str := readFile("archivo.txt") //Leemos el archivo y guardamos en string
	fmt.Println(str)
	fmt.Println("***************************************")

	// SLICE
	slice := []byte(str) //Convertimos string a byte
	fmt.Println("***************************************")

	bits := byteToBits(slice, blockSize) //convertimos bytes a bits
	encode := aplicandoHamming(bits, blockSize, parityBits, infoBits)

	ascii := binToASCII(encode) //convertimos el resultado a texto
	fmt.Println(ascii)          // y lo mostramos

	_ = writeFile("codificado.txt", ascii) //EScribimos el resultado en el archivo

	fmt.Println("***************************************")
	fmt.Println("Decodificacion:")
	decode := decodeHamming(encode, blockSize, infoBits)
	asciiDeco := bitsToByte(decode)
	decoded := string(asciiDeco)
	fmt.Println(decoded)
	_ = writeFile("decodificado.txt", decoded[:len(str)])

}
