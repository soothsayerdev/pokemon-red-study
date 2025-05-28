package trainer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPokemon struct {
	name  string
	level int
}

func (m *mockPokemon) Attack() {}

func (m *mockPokemon) Name() string {
	return m.name
}

func (m *mockPokemon) Level() int {
	return m.level
}

func TestBrock_CaptureAndListPokemons(t *testing.T) {
	// Cria instancia
	brock := NewBrock("Brock")

	// Cria 2 mocks de pokemon
	p1 := &mockPokemon{name: "Geodude", level: 12}
	p2 := &mockPokemon{name: "Onix", level: 14}

	// Captura os mocks pokemons
	brock.CapturePokemon(p1)
	brock.CapturePokemon(p2)


	// Verifica se o nome do treinador esta correto
	assert.Equal(t, "Brock", brock.Name())

	// Lista os pokemons 
	pokemons := brock.ListPokemons()

	// Verifica se a quantidade esta correta
	assert.Len(t, pokemons, 2)

	// Verifica se os dados do primeiro pokemon estao corretos
	assert.Equal(t, "Geodude", pokemons[0].Name)
	assert.Equal(t, 12, pokemons[0].Level)

	// Verifica se os dados do segundo pokemon estao corretos.
	assert.Equal(t, "Onix", pokemons[1].Name)
	assert.Equal(t, 14, pokemons[1].Level)
}

func TestMisty_CaptureAndListPokemons(t *testing.T) {
	// Instancia de misty 
	misty := NewMisty("Misty")

	// Cria dois pokemons mockados
	p1 := &mockPokemon{name: "Staryu", level: 18}
	p2 := &mockPokemon{name: "Starmie", level: 21}

	// Captura os mocks
	misty.CapturePokemon(p1)
	misty.CapturePokemon(p2)

	// Verifica se o nome esta certo
	assert.Equal(t, "Misty", misty.Name())

	// Lista os mocks capturados
	pokemons := misty.ListPokemons()

	// Verifica se a quantidade esta certa
	assert.Len(t, pokemons, 2)

	// Verifica se os dados do pokemons estao certos
	assert.Equal(t, "Staryu", pokemons[0].Name)
	assert.Equal(t, 18, pokemons[0].Level)

	assert.Equal(t, "Starmie", pokemons[1].Name)
	assert.Equal(t, 21, pokemons[1].Level)
}

func TestMisty_ListPokemonsMoreEff(t *testing.T) {
	// instancia de misty
	misty := NewMisty("Misty")

	// Cria e captura os mocks pokemons
	p1 := &mockPokemon{name: "Staryu", level: 18}
	p2 := &mockPokemon{name: "Starmie", level: 21}

	misty.CapturePokemon(p1)
	misty.CapturePokemon(p2)

	// Testa o metodo ListPokemonsMoreEff
	pokemonsMoreEff := misty.ListPokemonsMoreEff()
	assert.Len(t, pokemonsMoreEff, 2)
	assert.Equal(t, "Staryu", pokemonsMoreEff[0].Name)
	assert.Equal(t, 18, pokemonsMoreEff[0].Level)

	assert.Equal(t, "Starmie", pokemonsMoreEff[1].Name)
	assert.Equal(t, 21, pokemonsMoreEff[1].Level)


}

func TestMisty_ListPokemonsNoEff(t *testing.T) {
	// instancia de misty
	misty := NewMisty("Misty")

	// Cria e captura os mocks pokemons
	p1 := &mockPokemon{name: "Staryu", level: 18}
	p2 := &mockPokemon{name: "Starmie", level: 21}

	misty.CapturePokemon(p1)
	misty.CapturePokemon(p2)

	// Testa o metodo ListPokemonsMoreEff
	pokemonsMoreEff := misty.ListPokemonsNoEff()
	assert.Len(t, pokemonsMoreEff, 2)
	assert.Equal(t, "Staryu", pokemonsMoreEff[0].Name)
	assert.Equal(t, 18, pokemonsMoreEff[0].Level)

	assert.Equal(t, "Starmie", pokemonsMoreEff[1].Name)
	assert.Equal(t, 21, pokemonsMoreEff[1].Level)

	
}


func TestMisty_ListPokemons(t *testing.T) {
	// instancia de misty
	misty := NewMisty("Misty")

	// Cria e captura os mocks pokemons
	p1 := &mockPokemon{name: "Staryu", level: 18}
	p2 := &mockPokemon{name: "Starmie", level: 21}

	misty.CapturePokemon(p1)
	misty.CapturePokemon(p2)

	// Testa o metodo ListPokemonsMoreEff
	pokemonsMoreEff := misty.ListPokemons()
	assert.Len(t, pokemonsMoreEff, 2)
	assert.Equal(t, "Staryu", pokemonsMoreEff[0].Name)
	assert.Equal(t, 18, pokemonsMoreEff[0].Level)

	assert.Equal(t, "Starmie", pokemonsMoreEff[1].Name)
	assert.Equal(t, 21, pokemonsMoreEff[1].Level)

	
}
