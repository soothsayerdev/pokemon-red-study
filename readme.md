# Pokemon Red Study Project

## Descrição

Este projeto é um exemplo prático criado para estudar conceitos fundamentais de design de software usando a linguagem Go.
O domínio do projeto é inspirado no universo Pokémon, modelando a estrutura de um ginásio (Gym), um treinador (Trainer) e pokémons (Pokemon).

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
