package gym

import "pokemon-red-study/internal/model"

// Pewter representa o Ginásio de Pewter, contendo informações sobre seu nome, tipo
// e o treinador responsável. O campo 'trainer' é uma abstração, permitindo que qualquer
// implementação da interface trainer seja utilizada, promovendo a injeção de dependência.
type Pewter struct {
	trainer trainer // dependência injetada via interface, desacoplada de implementações concretas
	name    string
	Type    string // intuitar expecificar
}

// NewPewterGym é um construtor que cria e retorna uma nova instância de Pewter Gym.
// O parâmetro 'trainer' é uma dependência injetada como interface, permitindo que
// qualquer objeto que implemente a interface seja utilizado.
// Isso facilita a flexibilidade, testabilidade e manutenção do código.
func NewPewterGym(name string, trainer trainer) *Pewter {
	return &Pewter{
		name:    name,
		Type:    "Pewter",
		trainer: trainer,
	}
}

// Name retorna o nome do ginásio.
// Embora simples, este método segue o princípio de encapsulamento,
// permitindo que informações sejam expostas de maneira controlada.
func (p *Pewter) Name() string {
	return p.name
}

// TrainerName retorna o nome do treinador responsável pelo ginásio.
// Esse método depende apenas da interface trainer, não da implementação concreta (ex.: Brock),
// o que caracteriza a inversão de dependência: o módulo de alto nível (Pewter)
// depende de uma abstração e não de uma implementação específica.
func (p *Pewter) TrainerName() string {
	return p.trainer.Name()
}

// ListGymsPokemons lista todos os pokémons do treinador do ginásio.
// Assim como TrainerName, este método opera exclusivamente sobre a abstração trainer, demonstrando a aplicação prática da injeção de dependência via interface.
func (p *Pewter) ListGymsPokemons() []model.PokemonDTO {
	return p.trainer.ListPokemons()
}
