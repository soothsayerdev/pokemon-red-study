package model
 

// Modelos: Definem como os dados são recebidos e enviados pela API.
type PokemonDTO struct {
	Name string
	Level int
}

type PokemonMove struct {
	Name string `json:"name"`
	LevelLearnedAt int `json:"level_learned_at"`
}

type PokemonResponse struct {
	Name string `json:"name"`
	Moves []PokemonMove `json:"moves.omitempty"`
}

type PokemonRequest struct {
	Name string	`query:"name" validate:"required"`
	Level string `query:"level" validate:"required"`
}