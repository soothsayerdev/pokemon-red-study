package handlers

import (
	"net/http"
	"pokemon-red-study/internal/model"

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
	err := (&echo.DefaultBinder{}).BindQueryParams(c, &payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	// Chama o serviço que consome a PokeAPI e retorna os movimentos filtrados

	// Retorna a resposta em JSON
	return c.JSON(http.StatusOK, model.PokemonResponse{
		Name: payload.Name,
	})
}

func GetGyms(c echo.Context)
