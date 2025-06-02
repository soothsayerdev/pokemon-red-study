package services

import (
	"fmt"
	"pokemon-red-study/internal/model"
)

// GetGymInfo retorna informações do gym pelo name
func GetGymInfo(name string) (model.GymResponse, error) {
	switch name {
	case "Pewter":
		return model.GymResponse{
			GymName:     "Pewter",
			TrainerName: "Brock",
			Pokemons: []model.GymPokemons{
				{Name: "Geodude", Level: 12},
				{Name: "Onix", Level: 14},
			},
		}, nil
	case "Cerulean":
		return model.GymResponse{
			GymName:     "Cerulean",
			TrainerName: "Misty",
			Pokemons: []model.GymPokemons{
				{Name: "Staryu", Level: 18},
				{Name: "Starmie", Level: 21},
			},
		}, nil
	default:
		return model.GymResponse{}, fmt.Errorf("Ginasio nao encontrado")
	}
}
