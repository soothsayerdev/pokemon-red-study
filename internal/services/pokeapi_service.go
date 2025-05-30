package services

import (
	"fmt"
	"pokemon-red-study/internal/model"

	"github.com/go-resty/resty/v2"
)

type pokeAPIMove struct {
	Move struct {
		Name string `json:"name"`
	} `json:"move"`
	VersionGroupDetail []struct {
		LevelLearnedAt  int `json:"level_learned_at"`
		MoveLearnMethod struct {
			Name string `json:"name"`
		} `json:"move_learn_method"`
	} `json:"version_group_details"`
}

type pokeAPIResponse struct {
	Name  string        `json:"name"`
	Moves []pokeAPIMove `json:"moves"`
}

// Busca os moves aprendidos por level até o nivel informado
func GetPokemonMovesFromAPI(name string, level int) (model.PokemonResponse, error) {
	client := resty.New()
	var pokeResp pokeAPIResponse

	resp, err := client.R().SetResult(&pokeResp).Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil {
		return model.PokemonResponse{}, fmt.Errorf("erro ao consultar a PokeAPI: %v", err)
	}
	if resp.IsError() {
		return model.PokemonResponse{}, fmt.Errorf("Pokemon nao encontrado na PokeAPI")
	}

	// Filtra os moves aprendidos por level até o nivel atual
	var moves []model.PokemonMove
	for _, m := range pokeResp.Moves {
		for _, vgd := range m.VersionGroupDetail {
			if vgd.MoveLearnMethod.Name == "level-up" && vgd.LevelLearnedAt <= level {
				moves = append(moves, model.PokemonMove{
					Name:           m.Move.Name,
					LevelLearnedAt: vgd.LevelLearnedAt,
				})
				break // Evita duplicidade se tiver multiplos vgd (version_group_details)
			}
		}
	}

	return model.PokemonResponse{
		Name:  pokeResp.Name,
		Moves: moves,
	}, nil
}
