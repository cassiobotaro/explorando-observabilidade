# Dia 5 - Métricas - Histogramas

Aplicação que demonstra a utilização de histogramas, que são utilizados para medir distribuições de valores ao longo do tempo.

Alguns exemplos da utilização de Contadores:

- Tempo de resposta de APIs: Monitorar a latência de APIs RESTful, permitindo identificar APIs com desempenho lento e gargalos.
- Distribuição de tipos de erros: Visualizar a frequência de diferentes tipos de erros, ajudando a identificar os problemas mais comuns e suas causas.
- Tempo de processamento de eventos: Medir o tempo que o sistema leva para processar eventos, como transações ou mensagens, detectando gargalos na infraestrutura.

## Pré-requisitos

- [Go](https://go.dev)
- [Jq](https://jqlang.github.io/jq/)

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
go run .
```

Para acessar o servidor:

```sh
curl localhost:8080
```

A resposta não possui conteúdo, e pode levar até 10 segundos.

## Notas:

Para filtrar a saída e observar o valores gravados distribuídos em diferentes baldes(buckets) utilizamos o comando abaixo:

```sh
 go run main.go | jq  '.ScopeMetrics[].Metrics[].Data.DataPoints[0] | {Bounds, BucketCounts, Count}'
```

Os limites(Bounds) de cada balde(BucketCounts) do histograma são incrementados conforme mais registros são adicionados à distribuição.

Além disto temos um contador(Count) do número total de elementos adicionados a distribuição.

É possível definir os limites do histograma com a opção `metric.WithExplicitBucketBoundaries()`.
