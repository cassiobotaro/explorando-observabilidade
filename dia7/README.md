# Dia 7 - Métricas - Contadores UpDown Observáveis (Async)

Aplicação que demonstra a utilização de contadores UpDown de forma assíncrona.

Os contadores UpDown observáveis permitem monitorar valores que aumentam e diminuem sem necessidade de instrumentação manual em cada mudança. O sistema de métricas consulta periodicamente o valor atual através de callbacks.

Alguns exemplos da utilização de Contadores UpDown Observáveis:

- **Tamanho atual de uma fila de mensagens**: Consulta periodicamente o número de mensagens pendentes em uma fila (Redis, RabbitMQ, etc.) sem precisar instrumentar cada operação de enqueue/dequeue.
- **Número de conexões ativas em um pool**: Monitora quantas conexões estão ativas no pool de banco de dados através de consultas periódicas ao pool, sem overhead em cada uso de conexão.
- **Uso atual de disco**: Consulta periodicamente o espaço utilizado no sistema de arquivos através de APIs do sistema operacional, reportando o valor atual em bytes.
- **Items em cache**: Monitora o número de itens armazenados em cache consultando periodicamente o tamanho do cache, sem precisar instrumentar cada operação de get/set.
- **Tarefas pendentes em um executor**: Verifica periodicamente quantas tarefas estão aguardando processamento em um ThreadPoolExecutor ou ProcessPoolExecutor.

A principal diferença entre contadores UpDown síncronos (dia 4) e observáveis é que os observáveis consultam o estado atual do sistema periodicamente, enquanto os síncronos requerem instrumentação manual para cada incremento/decremento. Isso é especialmente útil quando o estado é mantido por componentes externos ou quando a instrumentação manual seria muito custosa.

Os callbacks dos contadores UpDown observáveis são invocados automaticamente pelo metric reader no momento da coleta, permitindo que o sistema de métricas controle a frequência de observação sem impactar a performance da aplicação.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

Para adicionar tarefas à fila:

```sh
curl localhost:8080/adicionar_tarefa
```

Para processar tarefas da fila:

```sh
curl localhost:8080/processar_tarefa
```

## Notas:

Para filtrar a saída e observar o tamanho atual da fila de tarefas:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "fila.tarefas.pendentes") | .data.data_points[] | {"Tarefas Pendentes": .value}'
```

Observe que o valor é atualizado periodicamente sem necessidade de instrumentação manual em cada operação da fila. O callback consulta o tamanho real da fila a cada intervalo de exportação.

Diferente do dia 4 onde incrementávamos/decrementávamos manualmente, aqui apenas consultamos o estado atual do sistema, delegando ao OpenTelemetry a responsabilidade de coletar os valores periodicamente.
