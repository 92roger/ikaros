package services

import (
	"encoding/json"
	"ikaros/backend/model"
	"ikaros/backend/pkg/processamento"
	"net/http"
)

// ProcessarCSVHandler é o handler para processar o arquivo CSV
func ProcessarCSVHandler(w http.ResponseWriter, r *http.Request) {
	arquivoCSV := "./data/exemplo.csv" // Ajuste o caminho conforme necessário

	dados, err := model.LerDadosCSV(arquivoCSV)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dadosProcessados, err := processamento.ProcessarDados(dados)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dadosProcessados)
}
