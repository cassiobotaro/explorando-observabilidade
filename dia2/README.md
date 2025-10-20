# Dia 2 - Métrica personalizada

Uma aplicação que define métrica personalizada.

Diferente do dia 1 que utilizava apenas métricas padrões da sdk `opentelemetry-instrumentation-fastapi`, essa aplicação define uma métrica personalizada que conta o número de lançamentos de dados por valor.

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

A métrica personalizada `dado.rolagens` quantifica o número de rolagens(lançamentos) por atributo `dado.valor` que representa o valor de rolagem.

Através do comando abaixo, filtramos a saída para ficar mais legível o valor e o número de vezes que foi rolado.

```sh
python app.py | jq '.resource_metrics[].scope_metrics[].metrics[].data.data_points[] | {rolagem: .attributes."dado.valor", vezes:
.value}'
```
