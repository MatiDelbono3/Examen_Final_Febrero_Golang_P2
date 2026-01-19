package Dtos

type PublicacionRequest struct {
	Titulo    string `json:"titulo"`
	Contenido string `json:"contenido"`
	Autor     string `json:"autor"`
	Categoria string `json:"categoria"`
	Estado    string `json:"estado"`
}
