package handlers

import (
	"examen_final_febrero_golang_P2/Dtos"
	Services "examen_final_febrero_golang_P2/Services"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublicacionHandler struct {
	service Services.PublicacionServiceInterface
}

func NewPublicacionHandler(service Services.PublicacionServiceInterface) *PublicacionHandler {
	return &PublicacionHandler{
		service: service,
	}
}

func (handler *PublicacionHandler) Crear(c *gin.Context) {
	var request Dtos.PublicacionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	avion, err := handler.service.Crear(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, avion)
}
func (handlers *PublicacionHandler) ListarPaginado(c *gin.Context) {

	resp, err := handlers.service.ListarPaginado()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (handlers *PublicacionHandler) FiltrarPorCampoDinamico(c *gin.Context) {
	var request Dtos.FiltroCampoDinamicoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := handlers.service.FiltrarPorCampoDinamico(request.Campo, request.Valor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (handlers *PublicacionHandler) FiltrarPublicacionesActivas(c *gin.Context) {
	var request Dtos.PublicacionesActivasRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := handlers.service.FiltrarPublicacionesActivas(request.Estado)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (handlers *PublicacionHandler) BorrarPublicacion(c *gin.Context) {
	id := c.Param("id")
	idStr, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = handlers.service.BorrarPublicacion(idStr.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
