package main

import (
	"fmt"
	"pokemon-red-study/internal/gym"
	"pokemon-red-study/internal/pokemon"
	"pokemon-red-study/internal/trainer"
)

func main() {

	geodude := pokemon.NewGeodude(12)
	onix := pokemon.NewOnix(14)

	brock := trainer.NewBrock("Brock")

	brock.CapturePokemon(geodude)
	brock.CapturePokemon(onix)

	pewterGym := gym.NewPewterGym("Pewter", brock)

	fmt.Println("Pewter Gym Trainer: ", pewterGym.TrainerName())

	pewterGym.ListGymsPokemons()
}
