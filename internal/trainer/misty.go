package trainer

import "pokemon-red-study/internal/model"

type Misty struct {
	name     string
	pokemons []pokemon
}

func NewMisty(name string) *Misty {
	return &Misty{
		name: name,
	}
}

func (m *Misty) Name() string {
	return m.name
}

func (m *Misty) CapturePokemon(pokemon pokemon) {
	m.pokemons = append(m.pokemons, pokemon)
}

func (m *Misty) ListPokemons() []model.Pokemon {
	var pokemons []model.Pokemon
	for _, pokemon := range m.pokemons {
		pokemons = append(pokemons, model.Pokemon{
			Name: pokemon.Name(),
			Level: pokemon.Level(),
		})
	}
	return pokemons
}