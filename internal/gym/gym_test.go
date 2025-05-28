package gym

import (
	"pokemon-red-study/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockTrainer struct {
	name     string
	pokemons []model.Pokemon
}

func (m *MockTrainer) Name() string {
	return m.name
}

func (m *MockTrainer) ListPokemons() []model.Pokemon {
	return m.pokemons
}

func TestCeruleanGym(t *testing.T) {
	mockTrainer := &MockTrainer{
		name: "Misty",
		pokemons: []model.Pokemon{
			{Name: "Staryu"},
			{Name: "Starmie"},
		},
	}

	gym := NewCeruleanGym("Cerulean", mockTrainer)

	assert.Equal(t, "Cerulean", gym.Name())
	assert.Equal(t, "Misty", gym.TrainerName())
	assert.Len(t, gym.ListGymsPokemons(), 2)
	assert.Equal(t, "Staryu", gym.ListGymsPokemons()[0].Name)
	assert.Equal(t, "Starmie", gym.ListGymsPokemons()[1].Name)
}

func TestPewterGym(t *testing.T) {
	mockTrainer := &MockTrainer{
		name: "Brock",
		pokemons: []model.Pokemon{
			{Name: "Geodude"},
			{Name: "Onix"},
		},
	}

	gym := NewPewterGym("Pewter", mockTrainer)

	assert.Equal(t, "Pewter", gym.Name())
	assert.Equal(t, "Brock", gym.TrainerName())
	assert.Len(t, gym.ListGymsPokemons(), 2)
	assert.Equal(t, "Geodude", gym.ListGymsPokemons()[0].Name)
	assert.Equal(t, "Onix", gym.ListGymsPokemons()[1].Name)
}
