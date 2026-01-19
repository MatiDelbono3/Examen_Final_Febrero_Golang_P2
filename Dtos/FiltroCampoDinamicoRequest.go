package dtos

type FiltroCampoDinamicoRequest struct {
	Campo string  `json:"campo"`
	Valor float64 `json:"valor"`
}
