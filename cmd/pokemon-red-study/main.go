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

	fmt.Printf("Welcome to the %s Gym!\n", pewterGym.Name())
	fmt.Printf("The trainer is: %s\n", pewterGym.TrainerName())
	fmt.Println("Pokemons in this gym:")

	// cerulean 
	fmt.Printf("Welcome to the %s Gym!\n", pewterGym.Name())
	fmt.Printf("The trainer is: %s\n", pewterGym.TrainerName())
	fmt.Println("Pokemons in this gym:")

	//TODO: Alterar esse retorno para ser um array de pokemon e iterar sobre eles, listando seus nomes
	pewterGym.ListGymsPokemons()
}
