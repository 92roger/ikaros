package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectMongoDB() {
	uri := "mongodb://localhost:27017"
	if uri == "" {
		log.Fatal("MONGODB_URI não definida")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Erro ao criar cliente MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar com MongoDB: %v", err)
	}

	// Teste a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao pingar MongoDB: %v", err)
	} else {
		log.Printf("Connected!")
	}

	DB = client
}

func CreateIndices() {
	colecao := DB.Database("bi_prod").Collection("cl_dados_cadastrais")

	// Criando um índice no campo 'nome'
	_, err := colecao.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{"nome": 1}, // 1 para ordem ascendente
		},
	)
	if err != nil {
		log.Fatal("Erro ao criar índice: ", err)
	}

	// Adicione a criação de índices para outras coleções aqui, se necessário
}
