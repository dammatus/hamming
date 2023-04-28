package modules

func decodeHamming(encoded []byte, blockSize int, infoSize int) []byte {
	decoded := make([]byte, 0)
	var decodedBlock = make([]byte, infoSize)
	for k := 0; k < len(encoded); k += blockSize {
		blockEncoded := encoded[k : k+blockSize]
		var j = 0
		for i := 0; i < len(blockEncoded); i++ {
			if !isPowerOfTwo(i + 1) {
				decodedBlock[j] = blockEncoded[i]
				j++
			}
		}
		decoded = append(decoded, decodedBlock...)
	}
	return decoded
}
