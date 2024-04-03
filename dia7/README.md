# Dia 7 - Métricas - Contadores UpDown Observáveis (Async)

Aplicação que demonstra a utilização de contadores de incremento e decremento de forma assíncrona.

A instrumentação assíncrona é útil em várias circunstâncias, por exemplo:

- Quando a atualização de um contador não é computacionalmente barata, e você não quer que o thread de execução atual espere pela medição;
- As observações precisam ocorrer em frequências não relacionadas à execução do programa (ou seja, elas não podem ser medidas com precisão quando vinculadas a um ciclo de vida de solicitação);

Alguns exemplos da utilização desta métrica:

- O número máximo permitido de conexões abertas ociosas;
- O número de bytes em uso em heap;
- Número de goroutines que existem atualmente;

## Pré-requisitos

- [Go](https://go.dev)
- [Jq](https://jqlang.github.io/jq/)

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
go run .
```

## Notas:

Para uma visão mais detalhada das métricas, utilize o comando abaixo:

```sh
 go run main.go | jq '.ScopeMetrics[].Metrics[]'
```

Neste exemplo os contadores permanecem em 0, pois não há incremento ou decremento.
