package main

import (
	"fmt"
	"pokemon-red-study/internal/gym"
	"pokemon-red-study/internal/handlers"
	"pokemon-red-study/internal/pokemon"
	"pokemon-red-study/internal/trainer"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// Rota para acessar Swagger - Essa interface possibilita a visualização de todos os endpoints da API, a definição dos parâmetros e os exemplos de resposta.
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//Endpoint busca moves do pokemon
	e.GET("/pokemon/moves", handlers.GetPokemonMoves)

	geodude := pokemon.NewGeodude(12)
	onix := pokemon.NewOnix(14)

	brock := trainer.NewBrock("Brock")

	brock.CapturePokemon(geodude)
	brock.CapturePokemon(onix)

	// pewter
	pewterGym := gym.NewPewterGym("Pewter", brock) // colocar na api

	fmt.Printf("Welcome to the %s Gym!\n", pewterGym.Name())
	fmt.Printf("The trainer is: %s\n", pewterGym.TrainerName())
	fmt.Println("Pokemons in this gym:")

	pokemons_brock := brock.ListPokemons()
	for _, p := range pokemons_brock {
		fmt.Printf("Name: %s, Level: %d\n", p.Name, p.Level)
	}

	staryu := pokemon.NewStaryu(18)
	starmie := pokemon.NewStarmie(21)

	misty := trainer.NewMisty("Misty")
	misty.CapturePokemon(staryu)
	misty.CapturePokemon(starmie)
	misty.CapturePokemon(geodude)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------------------------")
	fmt.Println()

	// cerulean
	ceruleanGym := gym.NewCeruleanGym("Cerulean", misty)

	fmt.Printf("Welcome to the %s Gym!\n", ceruleanGym.Name())
	fmt.Printf("The trainer is: %s\n", ceruleanGym.TrainerName())
	fmt.Println("Pokemons in this gym:")

	pokemons_misty := misty.ListPokemons()
	for _, p := range pokemons_misty {
		fmt.Printf("Name: %s, Level: %d\n", p.Name, p.Level)
	}

	//TODO: Alterar esse retorno para ser um array de pokemon e iterar sobre eles, listando seus nomes
	pewterGym.ListGymsPokemons()

	e.Logger.Fatal(e.Start(":8080"))
}
