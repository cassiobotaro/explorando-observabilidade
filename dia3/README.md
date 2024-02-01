# Dia 3 - Métricas - Contadores

Aplicação que demonstra a utilização de contadores como métricas.
Os contadores podem ser usados para medir um valor não negativo e crescente.

Alguns exemplos da utilização de Contadores:

- Número de erros (HTTP 500);
- Total de itens aguardando ser processados;
- Número de cache hit ou misses;
- Número de queries executadas em um banco de dados;
- Total de uso de CPU utilizado por um processo.

## Pré-requisitos

- [Go](https://go.dev)
- [Jq](https://jqlang.github.io/jq/)

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
go run .
```

Para acessar a rota de lançamentos de dados:

```sh
curl  localhost:8080/
```

ou

```sh
curl localhost:8080/rand
```

## Notas:

Contadores, como o nome diz, servem para contabilizar alguma coisa. Através do comando abaixo podemos ver o número de chamadas que a nossa aplicação obteve, através de uma métrica chamada `api.chamadas`.

```sh
go run . | jq '.ScopeMetrics[].Metrics[].Data.DataPoints[] | {Valor: .Value}'
```

Esta contagem pode ser agregada através da utilização de atributos. Será necessário chamar a rota `/rand` algumas vezes.

```sh

go run . | jq '.ScopeMetrics[].Metrics[].Data.DataPoints[] | {Valor: .Value, Chamadas: .Attributes[].Value.Value}'
```

Os contadores NÃO são somente incrementos unários(+1) e são utilizados para medir valores não negativos e crescentes.

No [dia 2](../dia2/) foi utilizado um contador para determinar o número de rolagens por valor do dado.
