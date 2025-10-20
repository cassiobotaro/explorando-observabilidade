# Dia 5 - Métricas - Histogramas

Aplicação que demonstra a utilização de histogramas, que são utilizados para medir distribuições de valores ao longo do tempo.

Alguns exemplos da utilização de Contadores:

- Tempo de resposta de APIs: Monitorar a latência de APIs RESTful, permitindo identificar APIs com desempenho lento e gargalos.
- Distribuição de tipos de erros: Visualizar a frequência de diferentes tipos de erros, ajudando a identificar os problemas mais comuns e suas causas.
- Tempo de processamento de eventos: Medir o tempo que o sistema leva para processar eventos, como transações ou mensagens, detectando gargalos na infraestrutura.

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

Para filtrar a saída e observar o valores gravados distribuídos em diferentes baldes(buckets) utilizamos o comando abaixo:

```sh
 python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "tarefa.duracao") | .data.data_points[] | { limites_buckets: .explicit_bounds, contagens: .bucket_counts }'
```

Os limites(explicit_bounds) de cada balde(bucket_counts) do histograma são incrementados conforme mais registros são adicionados à distribuição.

Além disto temos um contador(count) do número total de elementos adicionados a distribuição.

É possível definir os limites do histograma com a opção `explicit_bucket_boundaries_advisory`.
