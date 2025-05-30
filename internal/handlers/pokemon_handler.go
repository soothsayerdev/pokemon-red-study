package handlers

import (
	"net/http"
	"pokemon-red-study/internal/model"
	"pokemon-red-study/internal/services"

	"github.com/labstack/echo/v4"
)

// GetPokemonMoves godoc
// @Summary Obtém os movimentos de um Pokémon
// @Description Consulta a PokeAPI e retorna os movimentos aprendidos por level-up até um nível especificado.
// @Tags Pokémon
// @Accept json
// @Produce json
// @Param name query string true "Nome do Pokémon"
// @Param level query string true "Nível atual do Pokémon"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /pokemon/moves [get]
func GetPokemonMoves(c echo.Context) error { 
	payload := model.PokemonRequest{}
	// Faz o binding dos parametros da query para a struct
	err := (&echo.DefaultBinder{}).BindQueryParams(c, &payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if payload.Name == "" || payload.Level <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "nome e nivel são obrigatorios"})
	}

	// Chama o serviço que consome a PokeAPI e retorna os movimentos filtrados
	resp, err := services.GetPokemonMovesFromAPI(payload.Name, payload.Level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	// Retorna a resposta em JSON
	return c.JSON(http.StatusOK, resp)
	// return c.JSON(http.StatusOK, model.PokemonResponse{
	// 	Name: payload.Name,
	// })

}

func GetGyms(c echo.Context) error {
	payload := model.GymRequest{}
	err := (&echo.DefaultBinder{}).BindQueryParams(c, &payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, model.GymResponse{
		TrainerName: payload.Name,
		
	})
}
