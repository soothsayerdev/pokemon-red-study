package gym

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
