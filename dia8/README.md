# Dia 8 - Métricas - Gauges Observáveis (Async)

Aplicação que demonstra a utilização de gauges de forma assíncrona (observáveis).

Os gauges observáveis medem valores que podem subir ou descer, representando um **snapshot instantâneo** de um valor no momento da observação. Diferente dos contadores UpDown, gauges não são valores cumulativos - eles simplesmente reportam o valor atual de algo.

**Diferenças importantes:**

- **Observable Gauge (Dia 8)**: Mede um valor instantâneo não-aditivo (ex: temperatura atual é 25°C, não faz sentido somar temperaturas de diferentes momentos)
- **Observable UpDown Counter (Dia 7)**: Mede um valor cumulativo que pode aumentar/diminuir (ex: 5 conexões ativas agora, faz sentido somar/subtrair conexões)
- **Observable Counter (Dia 6)**: Mede um valor cumulativo crescente (ex: 3600 segundos de uptime, valor sempre cresce)

Alguns exemplos da utilização de Gauges Observáveis:

- **Uso de CPU (%)**: Consulta periodicamente o percentual de uso de CPU do processo. Representa o uso instantâneo, não cumulativo - 50% de CPU agora não se soma com 30% de antes.
- **Uso de memória (bytes)**: Monitora a quantidade de memória RAM utilizada pelo processo no momento. É um snapshot instantâneo do uso atual.
- **Temperatura do sistema**: Lê periodicamente a temperatura de componentes (CPU, GPU). Cada leitura é independente - 60°C agora substitui a leitura anterior.
- **Tamanho de um arquivo de log**: Consulta periodicamente o tamanho atual de um arquivo. O valor pode aumentar ou diminuir se o arquivo for truncado.
- **Latência da última requisição**: Registra quanto tempo levou a última operação executada. Não é cumulativo, apenas o valor mais recente.
- **Nível de bateria (%)**: Percentual de carga atual da bateria. Valor instantâneo que substitui o anterior a cada leitura.

A escolha entre Observable Gauge e Observable UpDown Counter depende da natureza do dado:
- Use **Gauge** se o valor é uma medida instantânea não-aditiva (temperatura, percentuais, níveis)
- Use **UpDown Counter** se o valor é aditivo/cumulativo (contadores, tamanhos de filas, conexões ativas)

Os callbacks são invocados periodicamente pelo metric reader, permitindo coletar métricas de recursos do sistema sem impactar a performance da aplicação.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

A aplicação irá coletar métricas do sistema automaticamente. Para gerar carga e observar mudanças nas métricas:

```sh
curl localhost:8080/trabalho_pesado
```

Para acessar a rota principal:

```sh
curl localhost:8080/
```

Para criar threads temporárias e observar mudanças na contagem:

```sh
curl localhost:8080/criar_threads
```

## Notas:

Para filtrar a saída e observar o uso de CPU:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "processo.cpu.uso") | .data.data_points[] | {"Uso de CPU (%)": .value}'
```

Para observar o uso de memória:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "processo.memoria.uso") | .data.data_points[] | {"Memória (MB)": .value}'
```

Para observar o número de threads:

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "processo.threads.contagem") | .data.data_points[] | {"Threads": .value}'
```

Observe que estes valores representam snapshots instantâneos do estado do sistema - cada leitura é independente e representa o estado atual, não valores acumulados ao longo do tempo.

**Por que usar Gauge em vez de UpDown Counter?**
- **Uso de CPU (%)**: É uma medida instantânea relativa, não um contador absoluto. 50% de CPU não é "50 unidades de algo contável".
- **Memória (MB)**: Embora seja um valor que pode aumentar/diminuir, representa um estado instantâneo não-aditivo do sistema.
- **Threads**: Embora pudesse ser um UpDown Counter (contagem de recursos), como Gauge representa melhor a ideia de "quantidade atual neste momento".
