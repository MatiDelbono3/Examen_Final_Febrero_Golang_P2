package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sucursal struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Titulo    string             `bson:"titulo"`
	Autor     string             `bson:"autor"`
	Contenido string             `bson:"contenido"`
	Categoria string             `bson:"categoria"`
	Estado    string             `bson:"estado"`
	CreadoEn  time.Time          `bson:"creadoEn"`
}
