package gym

import "pokemon-red-study/internal/model"

type trainer interface {
	Name() string
	ListPokemons() []model.PokemonDTO
}
