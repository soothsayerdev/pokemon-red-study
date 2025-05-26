package gym

type Pewter struct {
	trainer trainer
	name    string
	Type    string
}

func NewPewterGym(name string, trainer trainer) *Pewter {
	return &Pewter{
		name:    name,
		Type:    "Pewter",
		trainer: trainer,
	}
}

func (p *Pewter) Name() string {
	return p.name
}

func (p *Pewter) TrainerName() string {
	return p.trainer.Name()
}

func (p *Pewter) ListGymsPokemons() {
	p.trainer.ListPokemons()
}
