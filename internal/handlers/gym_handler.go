package handlers

import (
	"net/http"
	"pokemon-red-study/internal/model"
	"pokemon-red-study/internal/services"

	"github.com/labstack/echo/v4"
)

// GetGymInfo godoc
// @Summary Obtém informações do ginásio
// @Description Retorna o nome do treinador e os pokémons do ginásio informado.
// @Tags Ginásio
// @Accept json
// @Produce json
// @Param name query string true "Nome do Ginásio"
// @Success 200 {object} model.GymResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /gym [get]
func GetGymInfo(c echo.Context) error {
	payload := model.GymRequest{}
	// Faz o binding dos parametros da query para a struct GymRequest
	//err := (&echo.DefaultBinder{}).BindQueryParams(c, &payload)
	// if err != nil || payload.Name == "" {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nome do ginasio é obrigatorio"})
	// } required faz com q nao precise utilizar esse if

	// Chama a service para buscar as informações do ginasio
	resp, err := services.GetGymInfo(payload.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
