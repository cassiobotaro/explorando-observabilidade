import random
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
# O leitor de métricas periódico é vinculado com o(s) leitor(es) do fornecedor de métricas, sendo assim
# exportados (na console, no nosso caso).
fornecedor_medidores = MeterProvider(metric_readers=[leitor_metricas])

# registra o fornecedor de medidores criado como o global
metrics.set_meter_provider(fornecedor_medidores)

# poderia ser fornecedor_medidores.get_meter mas vamos assim pois ele é global
metrificador = metrics.get_meter("api")

# Como a execução será aleatória entre 0.5 e 6.5 segundos, vamos definir os seguintes baldes
# para distribuição explícita do histograma
limites = [
    0.5,  # 500ms
    1.0,  # 1s
    1.5,  # 1.5s
    2.0,  # 2s
    3.0,  # 3s
    4.0,  # 4s
    5.0,  # 5s
    6.0,  # 6s
    7.0,  # 7s (Cobre o máximo de 6.5s)
    7.5,  # 7.5s (Balde de segurança)
]
# histograma de duração de tarefa mede a duração de execução de uma tarefa em segundos
histograma_duracao_tarefa = metrificador.create_histogram(
    name="tarefa.duracao",
    description="A duração da execução da tarefa.",
    unit="s",
    explicit_bucket_boundaries_advisory=limites,
)

# Inicia uma aplicação web
app = FastAPI()


@app.get("/tarefa")
def executa_tarefa():
    # verifica o horário que a tarefa inicou a execução
    tempo_inicio = time.perf_counter()
    # simula a execução de uma tarefa com um delay
    # Mantendo o intervalo aleatório entre 0.5 e 6.5 segundos
    atraso_segundos = random.uniform(0.5, 6.5)
    time.sleep(atraso_segundos)

    # após a execução da tarefa, registra a duração no histograma
    tempo_fim = time.perf_counter()
    duracao = tempo_fim - tempo_inicio
    histograma_duracao_tarefa.record(
        duracao, attributes={"tarefa.tipo": "processamento_pesado"}
    )
    return {"mensagem": "Demorô um cadim!"}


if __name__ == "__main__":
    # Inicia o servidor web na porta 8080 com logs desligados para facilitar vizualização das métricas
    uvicorn.run(app, port=8080, log_config=None)
