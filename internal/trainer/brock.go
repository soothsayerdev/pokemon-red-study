package trainer

import "fmt"

type Brock struct {
	name     string
	pokemons []pokemon
}

func NewBrock(name string) *Brock {
	return &Brock{
		name: name,
	}
}

func (b *Brock) Name() string {
	return b.name
}

func (b *Brock) CapturePokemon(pokemon pokemon) {
	b.pokemons = append(b.pokemons, pokemon)
}

func (b *Brock) ListPokemons() {
	for _, pokemon := range b.pokemons {
		fmt.Println(pokemon.Name())
	}
}
