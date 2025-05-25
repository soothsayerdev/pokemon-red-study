package pokemon

type Geodude struct {
	name  string
	level int
}

func NewGeodude(level int) *Geodude {
	return &Geodude{
		level: level,
		name:  "Geodude",
	}
}

func (g *Geodude) Name() string {
	return g.name
}

func (g *Geodude) Attack() {
}
