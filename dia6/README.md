# Dia 6 - Métricas - Contadores Observáveis (Async)

Aplicação que demonstra a utilização de contadores de forma assíncrona.

A instrumentação assíncrona é útil em várias circunstâncias, por exemplo:

- Quando a atualização de um contador não é computacionalmente barata, e você não quer que o thread de execução atual espere pela medição;
- As observações precisam ocorrer em frequências não relacionadas à execução do programa (ou seja, elas não podem ser medidas com precisão quando vinculadas a um ciclo de vida de solicitação);

Alguns exemplos da utilização desta métrica:

- Tempo que a aplicação está de pé.
- Variações de CPU
- Tráfego de rede
- Tempo ocioso de CPU
- Tempo de execução de um processo

Se a cada variação de CPU gravarmos uma mudança em uma métrica, vamos ter um custo muito grande. Então, fazemos um modo assíncrono, onde de tempo em tempos fazemos a leitura do estado do que queremos medir.

## Pré-requisitos

- [Go](https://go.dev)
- [Jq](https://jqlang.github.io/jq/)

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
go run .
```

Como a métrica que olharemos é o uptime, nenhuma requisição será necessária.

## Notas:

Para filtrar a saída e observar o valores incrementados do tempo que a aplicação está rodando:

```sh
 go run main.go | jq '.ScopeMetrics[].Metrics[].Data.DataPoints[0] | {"Tempo de Atividade(s)": .Value}''
```
