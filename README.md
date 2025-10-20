# Explorando Observabilidade

🔭 Explorando e tomando notas sobre o mundo de observabilidade com Open Telemetry.

| Dia | Título                                                      |
| --: | :---------------------------------------------------------- |
|   1 | [otel + net/http](./dia1/)                                  |
|   2 | [Métrica personalizada](./dia2/)                            |
|   3 | [Métricas - Contadores](./dia3/)                            |
|   4 | [Métricas - Contadores UpDown](./dia4/)                     |
|   5 | [Métricas - Histogramas](./dia5/)                           |
|   6 | [Métricas - Contadores Observáveis (Async)](./dia6/)        |
|   7 | [Métricas - Contadores UpDown Observáveis (Async)](./dia7/) |

Extras:
- Recomendo assistir a [playlist](https://www.youtube.com/playlist?list=PLOQgLBuj2-3IL2SzHv1CHaBBHJEvHZE0m) sobre Observabilidade e Open Telemetry do [@dunossauro](https://github.com/dunossauro).

## Pré-requisitos
- [Python](https://www.python.org/)
- [uv](https://docs.astral.sh/uv/)
- [jq](https://jqlang.github.io/jq/)

Ideias futuras:

- Diferentes tipos de métricas
  - Observable (Async) Gauges
- Atributos em métricas
- Visões de métricas
- Adicionar trace
- Trace + complexo
- Como lidar com logs mesmo que sdk ainda não tenha ponte?
- Exportar métrica para alguma ferramenta / collector
- Exportar trace para alguma ferramenta / collector
- Exportar logs para alguma ferramenta / collector
- TelemetryGen e coletor otel
- Sampling
- Processors do coletor otel
- Instrumentção manual / automática
- coletor Load Balancer
- Testes de carga
- Tracetest
- logfire
- rastros de banco de dados
- rastros de clientes http

