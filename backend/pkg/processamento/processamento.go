package processamento

import (
	"ikaros/backend/model" // Importando o pacote model que contém as structs
	"strconv"
)

// ProcessarDados realiza o processamento dos dados lidos do CSV.
// Esta função recebe um slice de Dados e retorna um slice de DadosProcessados.
func ProcessarDados(dados []model.Dados) ([]model.DadosProcessados, error) {
	dados = RemoveDuplicidades(dados)

	var dadosProcessados []model.DadosProcessados
	for _, d := range dados {
		// Tratamento de campos nulos
		d.TratarCamposNulos()
		if d.Idade > 25 {
			processado := model.DadosProcessados{
				CodigoAcesso: GeraCodigoAcesso(), // Gera um código de acesso único
				Nome:         d.Nome,
				Idade:        d.Idade,
				Profissao:    d.Profissao,
				// Adicione lógica para outros campos
			}
			dadosProcessados = append(dadosProcessados, processado)
		}
	}

	return dadosProcessados, nil
}

func RemoveDuplicidades(dados []model.Dados) []model.Dados {
	unicos := make(map[string]model.Dados)
	var resultado []model.Dados

	for _, d := range dados {
		chave := d.Nome + "-" + strconv.Itoa(d.Idade) + "-" + d.Profissao
		if _, existe := unicos[chave]; !existe {
			unicos[chave] = d
			resultado = append(resultado, d)
		}
	}

	return resultado
}
