package gym

type Pewter struct {
	Name    string
	Type    string
	trainer trainer
}

type trainer interface {
	Name() string
	ListPokemons()
}

func NewPewterGym(name string, trainer trainer) *Pewter {
	return &Pewter{
		Name:    name,
		Type:    "Pewter",
		trainer: trainer,
	}
}

func (p *Pewter) TrainerName() string {
	return p.trainer.Name()
}

func (p *Pewter) ListGymsPokemons() {
	p.trainer.ListPokemons()
}
