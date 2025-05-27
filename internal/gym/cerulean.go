package gym

import "pokemon-red-study/internal/model"

// implementar gym misty

type Cerulean struct {
	trainer trainer
	name    string
	Type    string
}

func NewCeruleanGym(name string, trainer trainer) *Cerulean {
	return &Cerulean{
		name:    name,
		Type:    "Water",
		trainer: trainer,
	}
}

func (c *Cerulean) Name() string {
	return c.name
}

func (c *Cerulean) TrainerName() string {
	return c.trainer.Name()
}

func (c *Cerulean) ListGymsPokemons() []model.Pokemon {
	return c.trainer.ListPokemons()
}



