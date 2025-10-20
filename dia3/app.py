from random import randint

import uvicorn
from fastapi import FastAPI
from opentelemetry import metrics
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import (
    ConsoleMetricExporter,
    PeriodicExportingMetricReader,
)

# De quanto em quanto tempo as métricas serão exportadas
INTERVALO_EXPORTACAO_EM_SEGUNDOS = 10

# Cria um novo exportador de métricas para stdout (console, no nosso caso)
exportador_metricas = ConsoleMetricExporter()


# PeriodicExportingMetricReader é um leitor de métricas que coleta métricas
# em intervalos regulares e as exporta para o exportador atribuído.
leitor_metricas = PeriodicExportingMetricReader(
    # atribui um exportador ao leitor de métricas
    exporter=exportador_metricas,
    # Para fins demonstrativos, alteramos o intervalo de coleta
    # de métricas para alguns segundos.
    # O padrão é 1 minuto.
    export_interval_millis=INTERVALO_EXPORTACAO_EM_SEGUNDOS,
)

# Cria um novo fornecedor de medidores que é responsável por gerenciar
# os recursos necessários para coletar e exportar métricas
# O leitor de métricas periódico é vinculado com o(s) leitor(es) do forncedor de métricas, sendo assim
# exportados (na console, no nosso caso).
fornecedor_medidores = MeterProvider(metric_readers=[leitor_metricas])

# registra o fornecedor de medidores criado como o global
metrics.set_meter_provider(fornecedor_medidores)

# poderia ser fornecedor_medidores.get_meter mas vamos assim pois ele é global
metrificador = metrics.get_meter("api")

# contador_rolagens é uma métrica que conta o número de rolagens
# de dados por valor de rolagem
# Ex: 3 vezes o número 1, 2 vezes o número 2 etc
contador_rolagens = metrificador.create_counter(
    name="dado.rolagens",
    description="O número de rolagens por valor de rolagem",
    # A unidade de medida é definida como "{lancamento}" para que o valor
    # pluralizado seja exibido corretamente
    unit="{lancamento}",
)

# contador_chamadas é uma métrica que conta o número de chamadas à API
contador_chamadas = metrificador.create_counter(
    name="api.chamadas",
    description="O número de chamadas à API",
    # A unidade de medida é definida como "{chamada}" para que o valor
    # pluralizado seja exibido corretamente
    unit="{chamada}",
)


# Inicia uma aplicação web
app = FastAPI()


@app.get("/")
def raiz():
    # uma chamada é contada
    contador_chamadas.add(1)
    return "OK"


# Registra a rota /lançar_dado
@app.get("/lançar_dado")
def lançar_dado():
    # uma chamada é contada independentemente da rolagem do dado
    contador_chamadas.add(1)
    # A rolagem de dados é feita com um número aleatório entre 1 e 6
    rolagem = randint(1, 6)

    # Um atributo personalizado é criado com o valor da rolagem
    atributo_valor_rolagem = {"dado.valor": rolagem}
    # O atributo personalizado "dados.valor" é usado para agrupar as métricas
    # por valor de rolagem
    contador_rolagens.add(
        # O contador de rolagens é incrementado com o atributo personalizado
        1,
        attributes=atributo_valor_rolagem,
    )
    # o resultado é retornado
    return rolagem


if __name__ == "__main__":
    # Inicia o servidor web na porta 8080 com logs desligados para facilitar vizualização das métricas
    uvicorn.run(app, port=8080, log_config=None)
