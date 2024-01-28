package main

import (
	"bufio"                // Pacote para leitura de arquivos
	"encoding/csv"         // Pacote para manipulação de CSV
	"fmt"                  // Pacote para operações de entrada e saída
	"ikaros/backend/model" // Importe o pacote model, ajuste o caminho conforme necessário
	"ikaros/backend/pkg/processamento"
	"io"      // Pacote para operações de I/O, incluindo constantes como EOF
	"os"      // Pacote para operações do sistema operacional, como manipulação de arquivos
	"strconv" // Pacote para conversão de tipos de dados
)

func main() {
	// Abre o arquivo CSV
	file, err := os.Open("../../data/exemplo.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close() // Garante que o arquivo será fechado ao final da execução da função

	// Cria um leitor CSV a partir do arquivo
	reader := csv.NewReader(bufio.NewReader(file))

	// Variável para armazenar os dados lidos
	var dados []model.Dados

	// Lê o arquivo linha por linha
	for {
		linha, err := reader.Read()
		if err != nil {
			// Se chegou ao final do arquivo, sai do loop
			if err == io.EOF {
				break
			}
			// Imprime o erro, mas continua processando as outras linhas
			fmt.Println("Erro ao ler a linha:", err)
			continue
		}

		// Converte a idade de string para int
		idade, err := strconv.Atoi(linha[1])
		if err != nil {
			fmt.Println("Erro ao converter idade:", err)
			continue
		}

		// Cria um novo objeto Dados com as informações da linha e adiciona ao slice 'dados'
		dados = append(dados, model.Dados{
			Nome:      linha[0],
			Idade:     idade,
			Profissao: linha[2],
		})
	}

	// Chama a função ProcessarDados do pacote processamento
	dadosProcessados, err := processamento.ProcessarDados(dados)
	if err != nil {
		fmt.Println("Erro ao processar dados:", err)
		return
	}

	// Exibe os dados processados
	for _, d := range dadosProcessados {
		fmt.Printf("Cod. Acesso: %s, Nome: %s, Idade: %d, Profissao: %s\n", d.CodigoAcesso, d.Nome, d.Idade, d.Profissao)
		// Modifique para exibir outros campos de DadosProcessados conforme necessário
	}
}
