package services

import (
	"context"
	"encoding/json"
	"ikaros/backend/model"
	"ikaros/backend/pkg/database"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InserirDadosCadastraisHandler(w http.ResponseWriter, r *http.Request) {
	var dados model.DadosCadastrais

	// Decodificar o JSON do corpo da requisição para a estrutura de dados
	if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
		http.Error(w, "Entrada inválida: "+err.Error(), http.StatusBadRequest)
		return
	}

	colecao := database.DB.Database("bi_prod").Collection("cl_dados_cadastrais")

	// Verificar se já existe um registro com o mesmo 'nome'
	var resultado bson.M
	if err := colecao.FindOne(context.Background(), bson.M{"nome": dados.Nome}).Decode(&resultado); err != mongo.ErrNoDocuments {
		if err == nil {
			http.Error(w, "Registro já cadastrado", http.StatusBadRequest)
			return
		}
		http.Error(w, "Erro ao verificar duplicidade: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Inserir o novo registro, pois ele não existe na coleção
	_, err := colecao.InsertOne(context.Background(), dados)
	if err != nil {
		http.Error(w, "Erro ao inserir dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dados)
}
