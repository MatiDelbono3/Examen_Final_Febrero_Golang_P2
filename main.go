package main

import (
	"context"
	handlers "examen_final_febrero_golang_P2/Handlers"
	Service "examen_final_febrero_golang_P2/Services"
	"examen_final_febrero_golang_P2/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	// Mongo directo
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("examen").Collection("publicaciones")
	service := Service.NewPublicacionService(collection)

	handler := handlers.NewPublicacionHandler(service)

	r.POST("/publicacion", handler.Crear)
	r.GET("/publicacion", handler.ListarPaginado)
	r.GET("/publicacion/campo/valor", handler.FiltrarPorCampoDinamico)
	r.GET("/publicacion/estado", handler.FiltrarPublicacionesActivas)
	r.GET("/publicacion/id", handler.BorrarPublicacion)
	r.Run(":8080")
}
