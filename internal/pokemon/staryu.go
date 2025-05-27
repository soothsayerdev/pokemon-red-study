package pokemon

type Staryu struct {
	name string
	level int
}

func NewStaryu(level int) *Staryu {
	return &Staryu{
		level: level,
		name: "Staryu",
	}
}

func (s *Staryu) Name() string {
	return s.name
}

func (s *Staryu) Level() int {
	return s.level
}

func (s *Staryu) Attack() {
	
}