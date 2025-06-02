package model 

type GymRequest struct {
	Name string	`query:"name" validate:"required"`

}

type GymResponse struct {
	GymName	string 	`json:"gym_name"`
	TrainerName string `json:"trainerName"`
	GymType string `json:"gymType"`
	Pokemons []GymPokemons `json:"gymPokemons"`
}

type GymPokemons struct {
	Name string `json:"name"`
	Level int `json:"level"`
	Moves []string `json:"moves"`
} 