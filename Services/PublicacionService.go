package Service

import (
	"context"
	"errors"
	"time"

	"examen_final_febrero_golang_P2/Dtos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PublicacionServiceInterface interface {
	Crear(req Dtos.PublicacionRequest) (Dtos.PublicacionResponse, error)
	ListarPaginado(limit int, offset int) ([]Dtos.ListadoPaginacionResponse, error)
	FiltrarPorCampoDinamico(campo string, valor string) ([]Dtos.FiltroCampoDinamicoResponse, error)
	FiltrarPublicacionesActivas(estado string) ([]Dtos.PublicacionesActivasResponse, error)
	BorrarPublicacion(id string) error
}
type PublicacionService struct {
	collection     *mongo.Collection
	collectionName string
}

func NewPublicacionService(collection *mongo.Collection) *PublicacionService {
	return &PublicacionService{
		collection: collection,
	}

}
func (service *PublicacionService) Crear(req Dtos.PublicacionRequest) (Dtos.PublicacionResponse, error) {

	// Validaciones
	if req.Titulo == "" {
		return Dtos.PublicacionResponse{}, errors.New("el nombre es obligatorio")
	}

	if req.Autor == "" {
		return Dtos.PublicacionResponse{}, errors.New("latitud inválida")
	}

	// Documento a persistir
	doc := bson.M{
		"titulo":    req.Titulo,
		"autor":     req.Autor,
		"contenido": req.Contenido,
		"categoria": req.Categoria,
		"estado":    req.Estado,
		"creado_en": time.Now(),
	}

	result, err := service.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return Dtos.PublicacionResponse{}, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return Dtos.PublicacionResponse{
		Id:        id.Hex(),
		Titulo:    req.Titulo,
		Autor:     req.Autor,
		Categoria: req.Categoria,
		Contenido: req.Contenido,
		Estado:    req.Estado,
		CreadoEn:  time.Now(),
	}, nil
}
func (service *PublicacionService) ListarPaginado(limit int, offset int) ([]Dtos.ListadoPaginacionResponse, error) {

	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(offset))

	cursor, err := service.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var resultado []Dtos.ListadoPaginacionResponse

	for cursor.Next(context.Background()) {

		var doc struct {
			ID        primitive.ObjectID `bson:"_id"`
			Titulo    string             `bson:"titulo"`
			Autor     string             `bson:"autor"`
			Contenido string             `bson:"contenido"`
			Categoria string             `bson:"categoria"`
			Estado    string             `bson:"estado"`
			CreadoEn  time.Time          `bson:"creado_en"`
		}

		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		resultado = append(resultado, Dtos.ListadoPaginacionResponse{
			Titulo:    doc.Titulo,
			Autor:     doc.Autor,
			Contenido: doc.Contenido,
			Estado:    doc.Estado,
			CreadoEn:  doc.CreadoEn,
		})
	}

	return resultado, nil
}
func (service *PublicacionService) FiltrarPorCampoDinamico(campo string, valor string) ([]Dtos.FiltroCampoDinamicoResponse, error) {

	if campo == "" || valor == "" {
		return nil, errors.New("campo y valor son obligatorios")
	}

	collection := service.collection

	filter := bson.M{
		campo: valor,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var resultado []Dtos.FiltroCampoDinamicoResponse
	if err := cursor.All(context.TODO(), &resultado); err != nil {
		return nil, err
	}

	return resultado, nil
}
func (service *PublicacionService) FiltrarPublicacionesActivas(estado string) ([]Dtos.PublicacionesActivasResponse, error) {

	if estado == "" {
		return nil, errors.New("El estado es obligatorio")
	}

	collection := service.collection

	filter := bson.M{
		estado: estado,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var resultado []Dtos.PublicacionesActivasResponse
	if err := cursor.All(context.TODO(), &resultado); err != nil {
		return nil, err
	}

	return resultado, nil
}
func (service *PublicacionService) BorrarPublicacion(id string) error {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	_, err = service.collection.DeleteOne(
		context.Background(),
		bson.M{"_id": ObjectID},
	)
	return err
}
