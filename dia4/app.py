import time

import uvicorn
from fastapi import FastAPI
from opentelemetry import metrics
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import (
    ConsoleMetricExporter,
    PeriodicExportingMetricReader,
)

# De quanto em quanto tempo as métricas serão exportadas
INTERVALO_EXPORTACAO_MS = 10_000  # 10s

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
    export_interval_millis=INTERVALO_EXPORTACAO_MS,
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


# contador de requisições ativas mede o número de requisições HTTP ativas no momento
contador_requisições_ativas = metrificador.create_up_down_counter(
    name="http.servidor.requisicoes_ativas",
    description="Número de requisições HTTP ativas no momento",
    # A utilização de unit="1" é uma prática padrão no OpenTelemetry para métricas que representam contagens puras,
    # ou seja, valores que não possuem uma unidade física ou padrão de medida específico (como segundos, bytes, metros, etc.).
    unit="1",
)


# Inicia uma aplicação web
app = FastAPI()


@app.middleware("http")
async def middleware_de_requições_ativas(request, call_next):
    # ao iniciar a requisição, incrementa o contador de requisições ativas
    contador_requisições_ativas.add(
        1, attributes={"endpoint": request.url.path, "metodo": request.method}
    )
    # a requisição é processada
    resposta = await call_next(request)
    # ao finalizar a requisição, decrementa o contador de requisições ativas
    contador_requisições_ativas.add(
        -1, attributes={"endpoint": request.url.path, "metodo": request.method}
    )
    # por fim, retorna a resposta
    return resposta


@app.get("/lento")
def lento():
    # Simula um endpoint lento (10 segundos)
    time.sleep(10)
    return {"mensagem": "Demorô um cadim!"}


if __name__ == "__main__":
    # Inicia o servidor web na porta 8080 com logs desligados para facilitar vizualização das métricas
    uvicorn.run(app, port=8080, log_config=None)
