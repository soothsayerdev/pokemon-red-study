package pokemon

type Onix struct {
	name  string
	level int
}

func NewOnix(level int) *Onix {
	return &Onix{
		level: level,
		name:  "Onix",
	}
}

func (g *Onix) Name() string {
	return g.name
}

func (g *Onix) Attack() {
}
