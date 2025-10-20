# Dia 4 - Métricas - Contador UpDown

Aplicação que demonstra a utilização de contadores que podem incrementar e decrementar.

Os contadores UpDown podem ser incrementados e decrementados, permitindo que você observe um valor cumulativo que aumenta ou diminui.

Alguns exemplos da utilização de Contadores:

- Contagem de Conexões de Banco de Dados. Cada vez que uma nova conexão é estabelecida, o contador é incrementado; quando uma conexão é fechada, o contador é decrementado.
- Contagem de Threads Ativas em um Servidor Web. Quando uma nova requisição é recebida, o servidor pode iniciar uma nova thread, incrementando o contador. Quando uma thread completa o processamento de uma requisição e é finalizada, o contador é decrementado.
- Contagem de Clientes em uma Fila de Espera. Quando um cliente entra na fila, o contador é incrementado. Quando um cliente é atendido e sai da fila, o contador é decrementado.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

Para lançar 5 requisições sem aguardar seu retorno utilize o retorno:

```sh
for i in {1..5}; do curl  http://127.0.0.1:8080/lento & ; done
```

Como o endpoint é lento, será possível observar as requisições ativas no momento que a métrica é exportada.

## Notas:

Para filtrar a saída e observar o valor da métrica atualizando a medida que requisições são executadas, utilizamos o comando abaixo:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "http.servidor.requisicoes_ativas") | .data.data_points[] | {atributos: .attributes, valor: .value}'
```

Assim como no [dia 3](../dia3/) utilizamos contadores, mas dessa vez com a diferença de que podemos ter valores negativos e seu valor pode ser decrementado durante a execução do sistema.

Os contadores UpDown NÃO são somente incrementos ou decrementos unários(+1).
