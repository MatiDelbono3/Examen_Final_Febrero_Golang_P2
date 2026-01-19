package Dtos

import "time"

type PublicacionResponse struct {
	Id        string    `json:"id"`
	Titulo    string    `json:"titulo"`
	Contenido string    `json:"contenido"`
	Autor     string    `json:"autor"`
	Categoria string    `json:"categoria"`
	Estado    string    `json:"estado"`
	CreadoEn  time.Time `json:"creadoEn"`
}
