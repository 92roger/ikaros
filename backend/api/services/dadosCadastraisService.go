package services

import (
	"context"
	"encoding/json"
	"ikaros/backend/model"
	"ikaros/backend/pkg/database"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// ProcessarDadosCadastraisHandler é o handler para processar dados da coleção cl_dados_cadastrais
func ProcessarDadosCadastraisHandler(w http.ResponseWriter, r *http.Request) {
	colecao := database.DB.Database("bi_prod").Collection("cl_dados_cadastrais")
	ctx := context.Background()

	cursor, err := colecao.Find(ctx, bson.M{}) // bson.M{} é um filtro vazio que seleciona todos os documentos
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var dados []model.DadosCadastrais
	if err = cursor.All(ctx, &dados); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dados)
}
