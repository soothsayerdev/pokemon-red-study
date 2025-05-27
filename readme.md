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
- Organização de código em pacotes para melhor legibilidade e manutenção.
- Documentação através de Swagger.


### Debug
- Projeto já possui `launch.json` configurado para debug a partir do VSCode.

### CI
- Projeto possui integração com `GH Actions` para rodar testes automaticamente e build em cada `push` ou `pull request`, simulando uma Pipeline de CI/CD.

## Próximos Passos (fazer em ordem)

- [x] Implementar `Model` pokémon (contendo nome e level) para ser retornado como valor no `ListPokemons`. // 26/05
- [x] Adicionar o ginásio de `Cerulean` e a treinadora `Misty`. // 26/05
- [x] Implementar `mockery` para gerar mocks de interfaces e ler documentaçao do mockery https://vektra.github.io/mockery/latest/configuration/#mockery-init 27/05
- [ ] Implementar teste unitário (usando Testify) para o `Gym` e `Trainer`.
- [ ] Implementar chamada `HTTP` (utilizando `Resty`) para PokeAPI e buscar informações dos pokémons (com foco em `Habilidades` que ele sabe no nível atual)
- [ ] Implementar o restante dos ginários de `Kanto` (Elite Four não inclusa)
- [ ] Implementar Echo para criar uma API RESTful e retornar os resultados a partir de um GET e não pelo console
- [ ] Implementar Swagger para documentar a API RESTful

## Material de Estudo

- [Kanto Gym Leaders](https://pokemondb.net/red-blue/gymleaders-elitefour)
- [Mockery](https://vektra.github.io/mockery/latest/)
- [Testify](https://github.com/stretchr/testify)
- [Resty](https://github.com/go-resty/resty)
- [PokeAPI](https://pokeapi.co/)
- [Echo](https://echo.labstack.com/docs)
- [Swaggo](https://github.com/swaggo/echo-swagger)
- [JSON To Go](https://mholt.github.io/json-to-go/)


## Desafio

- [ ] Implementar Elite Four
