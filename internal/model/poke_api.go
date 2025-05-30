package model

// Modelos: Definem como os dados são recebidos e enviados pela API.
type Move struct {
	Move struct {
		Name string `json:"name"`
	} `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt int `json:"level_learned_at"`
	MoveLearnMethod struct {
		Name string `json:"name"`
	} `json:"move_learn_method"`
	VersionGroup struct {
		Name string `json:"name"`
	} `json:"version_group"`
}

type PokemonApi struct {
	Name  string `json:"name"`
	Moves []Move `json:"moves"`
}