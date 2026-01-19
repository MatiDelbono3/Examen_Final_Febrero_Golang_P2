package Dtos

type PublicacionesActivasResponse struct {
	Id        string `json:"id"`
	Titulo    string `json:"titulo"`
	Contenido string `json:"contenido"`
	Autor     string `json:"autor"`
	Categoria string `json:"categoria"`
}
