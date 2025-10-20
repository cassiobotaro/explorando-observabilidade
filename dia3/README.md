# Dia 3 - Métricas - Contadores

Aplicação que demonstra a utilização de contadores como métricas.
Os contadores podem ser usados para medir um valor não negativo e crescente.

Alguns exemplos da utilização de Contadores:

- Número de erros (HTTP 500);
- Total de itens aguardando ser processados;
- Número de cache hit ou misses;
- Número de queries executadas em um banco de dados;
- Total de uso de CPU utilizado por um processo.

## Como rodar

Para rodar a aplicação utilize o comando:

```sh
python app.py
```

Para acessar a rota principal:

```sh
curl  localhost:8080/
```

Para acessar a rota de lançamentos de dados:

```sh
curl  localhost:8080/lançar_dado
```

## Notas:

Ao acessar a raiz `/` ou `lançar_dado`, a aplicação incrementa um contador que mede o número de chamadas recebidas.

Contadores, como o nome diz, servem para contabilizar alguma coisa. Através do comando abaixo podemos ver o número de chamadas que a nossa aplicação obteve, através de uma métrica chamada `api.chamadas`.

```sh
 python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "api.chamadas") | {chamadas: .data.data_points [].value}'
```

Assim como no dia 2, a aplicação também define uma métrica personalizada que conta o número de lançamentos de dados por valor.
Esta contagem será agregada através da utilização de atributos. Será necessário chamar a rota `/lançar_dado` algumas vezes.

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[] | select(.name == "dado.rolagens") | {rolagem: .attributes."dado.valor", vezes: .value}'
```

Cuidado ao utilizar atributos pois pode gerar uma alta cardinalidade de métricas, ou seja, muitas combinações de atributos que podem impactar na performance do sistema de monitoramento.

Os contadores **NÃO** são somente incrementos unários(+1) e são utilizados para medir valores **não negativos** e **crescentes**.
