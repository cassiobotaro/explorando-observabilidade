# Dia 5 - Métricas - Histogramas

Aplicação que demonstra a utilização de histogramas, que são utilizados para medir distribuições de valores ao longo do tempo.

Alguns exemplos da utilização de Histogramas:

- **Tempo de resposta de APIs**: Registra a duração de cada requisição HTTP em buckets (ex: 0-100ms, 100-500ms, 500ms-1s, >1s). Permite calcular percentis (p50, p95, p99) e identificar requisições lentas que afetam a experiência do usuário.
- **Tamanho de requisições/respostas HTTP**: Mede o tamanho em bytes de requisições recebidas ou respostas enviadas, distribuindo em buckets. Útil para identificar payloads muito grandes que podem causar problemas de performance.
- **Tempo de processamento de eventos**: Registra quanto tempo cada evento leva para ser processado (ex: mensagens de fila, transações). Permite identificar gargalos e eventos que demoram mais que o esperado.
- **Duração de queries de banco de dados**: Mede o tempo de execução de queries SQL distribuído em buckets. Ajuda a identificar queries lentas e otimizar índices e consultas.
- **Latência de chamadas entre serviços**: Monitora o tempo de comunicação entre microserviços, permitindo detectar problemas de rede ou serviços com degradação de performance.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

Para acessar o servidor:

```sh
curl localhost:8080/tarefa
```

A resposta levará entre 0.5 a 6.5 segundos.

## Notas:

Para filtrar a saída e observar os valores gravados distribuídos em diferentes baldes(buckets) utilizamos o comando abaixo:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "tarefa.duracao") | .data.data_points[] | { limites_buckets: .explicit_bounds, contagens: .bucket_counts }'
```

Os limites(explicit_bounds) de cada balde(bucket_counts) do histograma são incrementados conforme mais registros são adicionados à distribuição.

Além disto temos um contador(count) do número total de elementos adicionados a distribuição.

É possível definir os limites do histograma com a opção `explicit_bucket_boundaries_advisory`.
