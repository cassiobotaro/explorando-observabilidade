# Dia 4 - Métricas - Contador UpDown

Aplicação que demonstra a utilização de contadores que podem incrementar e decrementar.

Os contadores UpDown podem ser incrementados e decrementados, permitindo que você observe um valor cumulativo que aumenta ou diminui.

Alguns exemplos da utilização de Contadores UpDown:

- **Contagem de Conexões de Banco de Dados**: Incrementado cada vez que uma nova conexão é estabelecida (+1) e decrementado quando uma conexão é fechada (-1). O valor atual representa o número de conexões ativas no momento.
- **Contagem de Threads Ativas em um Servidor Web**: Incrementado quando uma nova thread é iniciada para processar uma requisição (+1) e decrementado quando a thread finaliza o processamento (-1). Útil para monitorar a carga atual do servidor.
- **Contagem de Clientes em uma Fila de Espera**: Incrementado quando um cliente entra na fila (+1) e decrementado quando é atendido e sai da fila (-1). Permite observar o tamanho atual da fila em tempo real.
- **Uso de Memória em Bytes**: Incrementado quando memória é alocada (+N bytes) e decrementado quando é liberada (-N bytes). O valor representa a quantidade total de memória em uso pela aplicação.
- **Total de Itens Aguardando Processamento**: Incrementado quando novos itens são adicionados à fila de processamento e decrementado quando são processados. Permite identificar gargalos de processamento.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

Para lançar 5 requisições sem aguardar seu retorno utilize o comando:

```sh
for i in {1..5}; do curl  http://127.0.0.1:8080/lento & done
```

Como o endpoint é lento, será possível observar as requisições ativas no momento que a métrica é exportada.

## Notas:

Para filtrar a saída e observar o valor da métrica atualizando a medida que requisições são executadas, utilizamos o comando abaixo:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "http.servidor.requisicoes_ativas") | .data.data_points[] | {atributos: .attributes, valor: .value}'
```

Assim como no [dia 3](../dia3/) utilizamos contadores, mas dessa vez com a diferença de que podemos ter valores negativos e seu valor pode ser decrementado durante a execução do sistema.

Os contadores UpDown NÃO são somente incrementos ou decrementos unários (+1).
