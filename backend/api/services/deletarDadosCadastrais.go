package services

import (
	"context"
	"ikaros/backend/pkg/database"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeletarDadosCadastraisHandler lida com a "deleção" de dados cadastrais, realocando-os para outra coleção
func DeletarDadosCadastraisHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome") // Assumindo 'nome' como chave

	if nome == "" {
		http.Error(w, "Nome é necessário para deletar", http.StatusBadRequest)
		return
	}

	// Iniciar sessão para transação
	sessao, err := database.DB.StartSession()
	if err != nil {
		http.Error(w, "Erro ao iniciar sessão do MongoDB: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer sessao.EndSession(context.Background())

	// Executar operações dentro de uma transação
	err = mongo.WithSession(context.Background(), sessao, func(sc mongo.SessionContext) error {
		colecao := database.DB.Database("bi_prod").Collection("cl_dados_cadastrais")
		colecaoArquivada := database.DB.Database("bi_prod").Collection("cl_dados_cadastrais_apagados")

		// Encontrar o documento a ser "deletado"
		var documento bson.M
		if err := colecao.FindOne(sc, bson.M{"nome": nome}).Decode(&documento); err != nil {
			return err
		}

		// Inserir o documento na coleção de arquivados
		if _, err := colecaoArquivada.InsertOne(sc, documento); err != nil {
			return err
		}

		// Remover o documento da coleção original
		if _, err := colecao.DeleteOne(sc, bson.M{"nome": nome}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		http.Error(w, "Erro durante transação: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Dado realocado com sucesso"))
}
