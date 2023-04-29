package hamming

import (
	"fmt"
	"io/ioutil"

	"github.com/dammatus/hamming/modules"
)

const (
	bitsParity32    = 5
	bitsParity2048  = 11
	bitsParity65536 = 16
	bitsInfo32      = 26
	bitsInfo2048    = 2036
	bitsInfo65536   = 65519
)

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

	bits := modules.ByteToBits(slice, blockSize) //convertimos bytes a bits
	encode := modules.AplicandoHamming(bits, blockSize, parityBits, infoBits)

	ascii := modules.BinToASCII(encode) //convertimos el resultado a texto
	fmt.Println(ascii)                  // y lo mostramos

	_ = writeFile("codificado.txt", ascii) //EScribimos el resultado en el archivo

	fmt.Println("***************************************")
	fmt.Println("Decodificacion:")
	decode := modules.DecodeHamming(encode, blockSize, infoBits)
	asciiDeco := modules.BitsToByte(decode)
	decoded := string(asciiDeco)
	fmt.Println(decoded)
	_ = writeFile("decodificado.txt", decoded[:len(str)])

}
