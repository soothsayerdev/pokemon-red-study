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

func (m *Misty) ListPokemonsNoEff() []model.Pokemon {
	var pokemons []model.Pokemon 							// lento, pois poderá causar inumeras verificações e realocações desnecessesarias
	for _, pokemon := range m.pokemons {
		pokemons = append(pokemons, model.Pokemon{
			Name: pokemon.Name(),
			Level: pokemon.Level(),
		})
	}
	return pokemons
}

func (m *Misty) ListPokemons() []model.Pokemon {
	pokemons := make([]model.Pokemon, 0, len(m.pokemons)) // define capacidade inicial ( 2 pokemons na bag )
	for _, pokemon := range m.pokemons {
		pokemons = append(pokemons, model.Pokemon{			// rapido, pois nos ja definimos e cortamos o trabalho q a maquina teria de ler e entender, verificar e alocar
			Name: pokemon.Name(),
			Level: pokemon.Level(),
		})
	}
	return pokemons
}

func (m *Misty) ListPokemonsMoreEff() []model.Pokemon {
	pokemons := make([]model.Pokemon, len(m.pokemons))
	for i, pokemon := range m.pokemons { 	
		pokemons[i] = model.Pokemon{		// atribuição direta
			Name: pokemon.Name(),
			Level: pokemon.Level(),
		}
	}
	return pokemons
}