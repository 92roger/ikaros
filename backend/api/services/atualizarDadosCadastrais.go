package services

import (
	"context"
	"encoding/json"
	"ikaros/backend/model"
	"ikaros/backend/pkg/database"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// AtualizarDadosCadastraisHandler lida com a atualização de dados cadastrais
func AtualizarDadosCadastraisHandler(w http.ResponseWriter, r *http.Request) {
	var dados model.DadosCadastrais

	if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
		http.Error(w, "Entrada inválida: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Adicione mais validações aqui, se necessário

	colecao := database.DB.Database("bi_prod").Collection("cl_dados_cadastrais")
	_, err := colecao.UpdateOne(
		context.Background(),
		bson.M{"nome": dados.Nome}, // Substitua 'nome' pelo campo de identificação apropriado
		bson.M{"$set": dados},
	)
	if err != nil {
		http.Error(w, "Erro ao atualizar dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dados)
}
