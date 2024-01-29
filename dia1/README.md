# Dia 1 - otel + net/http

Uma aplicação utilizando o pacote padrão `net/http` com métricas fornecidas por padrão pela sdk de Go.

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

_Quais as métricas padrões monitoradas quando utilizo otelhttp?_

As trés métricas monitoradas são:

nome: http.server.request_content_length
descrição: Mede o tamanho do conteúdo da requisição HTTP (não compactado)

nome: http.server.response_content_length
descrição: Mede o tamanho do comprimento do conteúdo da resposta HTTP (não compactado)

nome: http.server.duration
descrição: Mede a duração do tratamento de requisições HTTP

```sh
go run . | jq '.ScopeMetrics[].Metrics[] | {Name: .Name, "Description": .Description}'
```

_Como ver uma métrica de forma indivdual_

Podemos selecionar a métrica através do seu nome assim evitando a saída mais poluída.

Exemplo:

```sh
go run . | jq '.ScopeMetrics[].Metrics[] | select(.Name == "http.server.duration")'
```

_O são esses atributos adicionados em cada métrica?_

Alguns atributos são adicionados a cada uma das métricas para ajudar em agregações e buscas.
Para vê-los, utilize o comando abaixo:

```sh
jq '.ScopeMetrics[0].Metrics[0].Data.DataPoints[0].Attributes'
```
