# Dia 6 - Métricas - Contadores Observáveis (Async)

Aplicação que demonstra a utilização de contadores de forma assíncrona.

A instrumentação assíncrona é útil em várias circunstâncias, por exemplo:

- Quando a atualização de um contador não é computacionalmente barata, e você não quer que o thread de execução atual espere pela medição;
- As observações precisam ocorrer em frequências não relacionadas à execução do programa (ou seja, elas não podem ser medidas com precisão quando vinculadas a um ciclo de vida de solicitação);

Alguns exemplos da utilização de Contadores Observáveis:

- **Tempo de atividade da aplicação (uptime)**: Mede há quanto tempo a aplicação está rodando em segundos. Coletado periodicamente através de um callback que calcula a diferença entre o tempo atual e o tempo de inicialização. Valor sempre crescente.
- **Total de bytes processados**: Soma cumulativa de todos os bytes lidos/escritos desde o início da aplicação, obtida periodicamente consultando estatísticas do processo.
- **Tempo total de CPU consumido**: Total acumulado de CPU time usado pelo processo desde seu início, consultado periodicamente através de APIs do sistema operacional.
- **Contador de page faults**: Número total de page faults ocorridos desde o início do processo, obtido periodicamente através de estatísticas do sistema.
- **Estatísticas de Garbage Collector**: Métricas como número total de ciclos de GC executados desde o início, tempo total acumulado de pausa, etc. Coletadas do runtime sem impacto na aplicação.

Observable Counters são apropriados para valores que **apenas crescem** ao longo do tempo. Se o valor pode aumentar e diminuir, use Observable UpDown Counter (Dia 7). Se o valor é uma medida instantânea não-cumulativa, use Observable Gauge (Dia 8).

Se a cada mudança de CPU gravarmos uma métrica de forma síncrona, o custo computacional seria muito alto. Com contadores observáveis (assíncronos), fazemos leituras periódicas do estado do sistema, reduzindo overhead e permitindo que o coletor de métricas controle a frequência de coleta.

Os callbacks dos contadores observáveis são invocados automaticamente pelo metric reader (como o `PeriodicExportingMetricReader`) no momento da coleta. A frequência de invocação depende da configuração do reader - no nosso caso, definida pelo parâmetro `export_interval_millis`.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

Como a métrica que olharemos é o uptime (tempo de atividade), nenhuma requisição será necessária.

## Notas:

Para filtrar a saída e observar os valores incrementados do tempo que a aplicação está rodando:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "app.atividade.tempo") | .data.data_points[] | {"Tempo de Atividade(s)": .value}'
```
