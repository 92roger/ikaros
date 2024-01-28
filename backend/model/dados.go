package model

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Dados representa uma linha do arquivo CSV.
type Dados struct {
	Nome      string
	Idade     int
	Profissao string
}

// TratarCamposNulos implementa a interface TratadorDeCampoNulo para Dados.
func (d *Dados) TratarCamposNulos() {
	if d.Nome == "" {
		d.Nome = "N/A"
	}
	if d.Profissao == "" {
		d.Profissao = "N/A"
	}
	// Repita para outros campos se necessário
}

// LerDadosCSV lê um arquivo CSV e retorna um slice de Dados.
func LerDadosCSV(caminhoArquivo string) ([]Dados, error) {
	// Abre o arquivo CSV
	file, err := os.Open(caminhoArquivo)
	if err != nil {
		return nil, err // Retorna o erro se não conseguir abrir o arquivo
	}
	defer file.Close() // Garante que o arquivo será fechado ao final da função

	// Cria um leitor CSV a partir do arquivo
	reader := csv.NewReader(bufio.NewReader(file))
	_, err = reader.Read() // Lê e ignora a primeira linha (cabeçalho)
	if err != nil {
		return nil, err
	}

	// Slice para armazenar os dados lidos
	var dados []Dados

	// Lê o arquivo linha por linha
	for {
		linha, err := reader.Read()
		if err != nil {
			if err == io.EOF { // Checa se chegou ao final do arquivo
				break
			}
			return nil, err // Retorna o erro se não conseguir ler a linha
		}

		// Parseia a idade de string para int
		idade, err := strconv.Atoi(linha[1])
		if err != nil {
			return nil, err // Retorna o erro se não conseguir converter a idade
		}

		// Adiciona os dados lidos ao slice
		dados = append(dados, Dados{
			Nome:      linha[0],
			Idade:     idade,
			Profissao: linha[2],
		})
	}

	// Retorna o slice de Dados e nil como erro, indicando sucesso
	return dados, nil
}

// DadosProcessados representa os dados após serem processados.
type DadosProcessados struct {
	CodigoAcesso string
	Nome         string
	Idade        int
	Profissao    string

	// Adicione outros campos que você deseja incluir após o processamento
	// Por exemplo: Categoria, Pontuação, etc.
}
