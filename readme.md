# Pokemon Red Study Project

## Descrição

Este projeto é um exemplo prático criado para estudar conceitos fundamentais de design de software usando a linguagem Go.
O domínio do projeto é inspirado no universo `Pokémon`, modelando a estrutura de um ginásio (`Gym`), um treinador (`Trainer`) e pokémons (`Pokemon`).

```bash
go run -race ./cmd/pokemon-red-study
Welcome to the Pewter Gym!
The trainer is: Brock
Pokemons in this gym:
Geodude
Onix
```

## O que estamos aprendendo?

### Injeção de Dependência (Dependency Injection)

- O `Gym` depende da interface `trainer`, e não de uma implementação concreta.

- O `trainer` é injetado no `Gym` via o *construtor* (`NewPewterGym`).

- Aprendemos como criar *dependências* mais flexíveis e testáveis usando `interfaces` e injeção via `construtor`.

### Inversão de Dependência (Dependency Inversion)

- O módulo `Gym` não depende diretamente de um módulo de baixo nível (`Brock`), mas sim de uma abstração (`trainer` interface).

- Aprendemos a `desacoplar` componentes, facilitando manutenção e evolução do código.

- Demonstramos como aplicar o princípio `D` do `SOLID` em projetos Go, de forma simples e idiomática.

### Boas Práticas

- Uso de interfaces para abstração.

- Encapsulamento de dados e comportamento.

- Separação de responsabilidades em pacotes: `gym`, `trainer` e `pokemon`.

- Como construir sistemas modulares e testáveis.


### Debug
- Projeto já possui `launch.json` configurado para debug a partir do VSCode.

## Próximos Passos

- [ ] Implementar `Model` pokémon (contendo nome e level) para ser retornado como valor no `ListPokemons`.
- [ ] Adicionar o ginário de `Cerulean` e a treinadora `Misty`.
- [ ] Implementar `mockery` para gerar mocks de interfaces.
- [ ] Implementar teste unitário para o `Gym` e `Trainer`.
- [ ] Implementar chamada `HTTP` para PokeAPI e buscar informações dos pokémons (com foco em `Habilidades` que ele sabe no nível atual)
- [ ] Implementar o restante dos ginários de `Kanto` (Elite Four não inclusa)
- [ ] Implementar Echo para criar uma API RESTful e retornar os resultados a partir de um GET e não pelo console

## Material de Estudo

- [Kanto Gym Leaders](https://pokemondb.net/red-blue/gymleaders-elitefour)
- [Echo](https://echo.labstack.com/docs)
- [Mockery](https://vektra.github.io/mockery/latest/)
- [PokeAPI](https://pokeapi.co/)


## Desafio

- [ ] Implementar Elite Four
