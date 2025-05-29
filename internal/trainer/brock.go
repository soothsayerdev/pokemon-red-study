package trainer

import (
	"pokemon-red-study/internal/model"
)

// Brock representa uma implementação concreta da interface trainer.
// Ele contém um nome e uma lista de pokémons capturados.
// Esta estrutura será injetada como dependência onde a interface trainer for requerida.
type Brock struct {
	name     string
	pokemons []pokemon // lista de pokémons capturados pelo treinador
}

// NewBrock é um construtor que cria e retorna uma nova instância de Brock.
// Esse construtor permite criar uma implementação concreta que será injetada onde for necessário um trainer, promovendo a injeção de dependência.
func NewBrock(name string) *Brock {
	return &Brock{
		name: name,
	}
}

// Name retorna o nome do treinador.
// Este método implementa parte do contrato definido pela interface trainer.
// A interface abstrai o comportamento do treinador, permitindo a inversão de dependência: módulos de alto nível dependem dessa interface, e não diretamente de Brock.
func (b *Brock) Name() string {
	return b.name
}

// CapturePokemon adiciona um Pokémon à lista de pokémons capturados pelo treinador.
// Essa função permite que o estado de Brock evolua dinamicamente.
func (b *Brock) CapturePokemon(pokemon pokemon) {
	b.pokemons = append(b.pokemons, pokemon)
}

// ListPokemons percorre e imprime no console o nome de todos os pokémons capturados.
// Este método também faz parte da interface trainer, permitindo que Brock seja injetado em qualquer componente que dependa da abstração, como o Gym.
// Isso facilita o desacoplamento e promove a injeção de dependência.
func (b *Brock) ListPokemons() []model.PokemonDTO {
	var pokemons []model.PokemonDTO
	for _, pokemon := range b.pokemons {
		pokemons = append(pokemons, model.PokemonDTO{
			Name:  pokemon.Name(),
			Level: pokemon.Level(),
		})
		// fmt.Println(pokemon.Name())

	}
	return pokemons
}
