from collections import deque

import uvicorn
from fastapi import FastAPI
from opentelemetry import metrics
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import (
    ConsoleMetricExporter,
    PeriodicExportingMetricReader,
)

# Fila de tarefas simulada
fila_tarefas = deque()

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
metrificador = metrics.get_meter("fila")


def observar_tamanho_fila(_):
    # Quando invocado, esta função retorna o tamanho atual da fila
    # o parâmetro de opções da observação não é utilizado aqui por isso _
    yield metrics.Observation(len(fila_tarefas))


# define uma métrica do tipo observable up-down counter para medir o tamanho da fila
metrificador.create_observable_up_down_counter(
    name="fila.tarefas.pendentes",
    # função chamada para observar o valor da métrica
    callbacks=[observar_tamanho_fila],
    description="Número de tarefas pendentes na fila",
    # A utilização de unit="1" é uma prática padrão no OpenTelemetry para métricas que representam contagens puras
    unit="1",
)


# Inicia uma aplicação web
app = FastAPI()


@app.get("/adicionar_tarefa")
def adicionar_tarefa():
    # Adiciona uma nova tarefa à fila
    tarefa_id = len(fila_tarefas) + 1
    fila_tarefas.append(f"tarefa-{tarefa_id}")
    return {
        "mensagem": f"Tarefa {tarefa_id} adicionada",
        "tarefas_pendentes": len(fila_tarefas),
    }


@app.get("/processar_tarefa")
def processar_tarefa():
    # Remove e processa uma tarefa da fila
    if fila_tarefas:
        tarefa = fila_tarefas.popleft()
        return {
            "mensagem": f"Tarefa {tarefa} processada",
            "tarefas_pendentes": len(fila_tarefas),
        }
    return {"mensagem": "Nenhuma tarefa pendente", "tarefas_pendentes": 0}


@app.get("/status")
def status():
    # Retorna o status atual da fila
    return {"tarefas_pendentes": len(fila_tarefas), "tarefas": list(fila_tarefas)}


if __name__ == "__main__":
    # Inicia o servidor web na porta 8080 com logs desligados para facilitar vizualização das métricas
    uvicorn.run(app, port=8080, log_config=None)
