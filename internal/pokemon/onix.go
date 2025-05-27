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

func (o *Onix) Name() string {
	return o.name
}

func (o *Onix) Attack() {
}

func (o *Onix) Level() int {
	return o.level
}