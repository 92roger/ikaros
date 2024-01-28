package processamento

import (
	"crypto/rand"
	"fmt"
)

// GeraCodigoAcesso cria um código de acesso único no formato 'IK-ABC1234'.
func GeraCodigoAcesso() string {
	const letras = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const tamanhoLetras = 3
	const tamanhoNumeros = 4

	// Gera letras aleatórias
	bLetras := make([]byte, tamanhoLetras)
	if _, err := rand.Read(bLetras); err != nil {
		panic(err) // Em um caso real, você pode querer tratar o erro de forma mais adequada
	}
	for i := 0; i < tamanhoLetras; i++ {
		bLetras[i] = letras[bLetras[i]%byte(len(letras))]
	}

	// Gera números aleatórios
	bNumeros := make([]byte, tamanhoNumeros)
	if _, err := rand.Read(bNumeros); err != nil {
		panic(err)
	}
	for i := 0; i < tamanhoNumeros; i++ {
		bNumeros[i] = '0' + (bNumeros[i] % 10)
	}

	return fmt.Sprintf("IK-%s%s", string(bLetras), string(bNumeros))
}
