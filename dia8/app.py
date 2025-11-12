import threading
import time

import psutil
import uvicorn
from fastapi import FastAPI
from opentelemetry import metrics
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import (
    ConsoleMetricExporter,
    PeriodicExportingMetricReader,
)

# Obtém o processo atual para monitoramento
processo_atual = psutil.Process()

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
metrificador = metrics.get_meter("sistema")


def observar_cpu(_):
    # Retorna o uso atual de CPU do processo em percentual
    # cpu_percent() retorna o uso desde a última chamada, com intervalo de 0.1s
    uso_cpu = processo_atual.cpu_percent(interval=0.1)
    yield metrics.Observation(uso_cpu)


def observar_memoria(_):
    # Retorna o uso atual de memória do processo em megabytes
    info_memoria = processo_atual.memory_info()
    memoria_mb = info_memoria.rss / 1024 / 1024  # Converte de bytes para MB
    yield metrics.Observation(memoria_mb)


def observar_threads(_):
    # Retorna o número de threads ativas do processo
    num_threads = processo_atual.num_threads()
    yield metrics.Observation(num_threads)


# Define gauge observável para uso de CPU
metrificador.create_observable_gauge(
    name="processo.cpu.uso",
    callbacks=[observar_cpu],
    description="Percentual de uso de CPU do processo",
    unit="%",
)

# Define gauge observável para uso de memória
metrificador.create_observable_gauge(
    name="processo.memoria.uso",
    callbacks=[observar_memoria],
    description="Uso de memória do processo",
    unit="MB",
)

# Define gauge observável para contagem de threads
metrificador.create_observable_gauge(
    name="processo.threads.contagem",
    callbacks=[observar_threads],
    description="Número de threads ativas do processo",
    unit="1",
)


# Inicia uma aplicação web
app = FastAPI()


@app.get("/")
def raiz():
    return {"mensagem": "Servidor monitorando métricas do sistema"}


@app.get("/trabalho_pesado")
def trabalho_pesado():
    # Simula trabalho computacional pesado para aumentar uso de CPU
    tempo_inicio = time.time()
    resultado = 0
    for i in range(10_000_000):
        resultado += i * i
    duracao = time.time() - tempo_inicio
    return {
        "mensagem": "Trabalho concluído",
        "duracao_segundos": round(duracao, 2),
        "resultado": resultado,
    }


@app.get("/criar_threads")
def criar_threads():
    # Cria algumas threads temporárias para demonstrar mudança na contagem
    def tarefa_thread():
        time.sleep(5)

    threads = []
    for i in range(3):
        thread = threading.Thread(target=tarefa_thread, daemon=True)
        thread.start()
        threads.append(thread)

    return {"mensagem": f"{len(threads)} threads criadas (duração: 5s)"}


if __name__ == "__main__":
    # Inicia o servidor web na porta 8080 com logs desligados para facilitar vizualização das métricas
    uvicorn.run(app, port=8080, log_config=None)
