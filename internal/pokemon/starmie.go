package pokemon

type Starmie struct {
	name string
	level int
}

func NewStarmie(level int) *Starmie{
	return &Starmie{
		level: level,
		name: "Starmie",
	}
}

func (s *Starmie) Name() string {
	return s.name
}

func (s *Starmie) Level() int {
	return s.level
}

func (s *Starmie) Attack() {

}