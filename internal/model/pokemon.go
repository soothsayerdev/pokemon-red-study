package model
 
type PokemonDTO struct {
	Name string
	Level int
}

type PokemonResponse struct {
	Name string `json:"name"`
	Moves []string `json:"moves,omitempty"`
}

type PokemonRequest struct {
	Name string	`query:"name" validate:"required"`
	Level string `query:"level" validate:"required"`
}