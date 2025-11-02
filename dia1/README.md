# Dia 1 - otel + fastapi

Uma aplicação utilizando `fastapi` com métricas fornecidas por padrão pela sdk `opentelemetry-instrumentation-fastapi`.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

Para acessar a rota de lançamentos de dados:

```sh
curl  localhost:8080/lançar_dado
```

## Notas:

_Quais as métricas padrões monitoradas?_

As trés métricas monitoradas são:

nome: http.server.active_requests
descrição: Número de solicitações ativas do servidor HTTP.

nome: http.server.duration
descrição: Mede a duração das solicitações HTTP de entrada.

nome: http.server.response.size
descrição: Mede o tamanho das mensagens de resposta HTTP (compactadas).

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | {nome: .name, descricao: .description}'
```

_Como ver uma métrica de forma individual_

Podemos selecionar a métrica através do seu nome assim evitando a saída mais poluída.

Exemplo:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "http.server.duration")'
```

_O que são esses atributos adicionados em cada métrica?_

Alguns atributos são adicionados a cada uma das métricas para ajudar em agregações e buscas.
Para vê-los, utilize o comando abaixo:

```sh
python app.py | jq '.resource_metrics[0].scope_metrics[0].metrics[0].data.data_points[0].attributes'
```
