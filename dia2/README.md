# Dia 2 - Métrica personalizada

Uma aplicação que define métrica personalizada utilizando sdk de Go.

## Pré-requisitos

- [Go](https://go.dev)
- [Jq](https://jqlang.github.io/jq/)

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
go run .
```

Para acessar a rota de lançamentos de dados:

```
curl  localhost:8080/lançardados
```

## Notas:

A métrica personalizada `dados.rolagens` quantifica o número de rolagens(lançamentos) por valor de rolagem.

```sh
go run . | jq '.ScopeMetrics[].Metrics[].Data.DataPoints[] | {Valor: .Value, Rolagens: .Attributes[].Value.Value}'
```

Através do comando acima, filtramos a saída para ficar mais legível o valor e o número de vezes que foi rolado.
