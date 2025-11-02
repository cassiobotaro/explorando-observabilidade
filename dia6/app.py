import time

import uvicorn
from fastapi import FastAPI
from opentelemetry import metrics
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import (
    ConsoleMetricExporter,
    PeriodicExportingMetricReader,
)

# Captura o tempo de início da aplicação
START_TIME = time.time()

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
metrificador = metrics.get_meter("app")


def observar_duracao_app(_):
    # Quando invocado, esta função retorna a duração desde o início da aplicação
    # o parâmetro de opções da observação não é utilizado aqui por isso _
    yield metrics.Observation(time.time() - START_TIME)


# define uma métrica do tipo observable gauge para medir a duração desde o início do aplicativo
metrificador.create_observable_gauge(
    name="app.atividade.tempo",
    # função chamada para observar o valor da métrica
    callbacks=[observar_duracao_app],
    description="A duração desde o início do aplicativo.",
    # A unidade de medida é segundos
    unit="s",
)


# Inicia uma aplicação web
app = FastAPI()


if __name__ == "__main__":
    # Inicia o servidor web na porta 8080 com logs desligados para facilitar vizualização das métricas
    uvicorn.run(app, port=8080, log_config=None)
