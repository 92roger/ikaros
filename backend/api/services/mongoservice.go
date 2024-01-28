package services

import (
	"context"
	"ikaros/backend/model"
	"ikaros/backend/pkg/database"
)

func CreateSomething(data model.DadosProcessados) error {
	collection := database.DB.Database("bi_prod").Collection("testesdatas")

	_, err := collection.InsertOne(context.Background(), data)
	return err
}
