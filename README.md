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
|   8 | [Métricas - Gauges Observáveis (Async)](./dia8/)            |

Extras:
- Recomendo assistir a [playlist](https://www.youtube.com/playlist?list=PLOQgLBuj2-3IL2SzHv1CHaBBHJEvHZE0m) sobre Observabilidade e Open Telemetry do [@dunossauro](https://github.com/dunossauro).

## Pré-requisitos

- [Python](https://www.python.org/) - Linguagem de programação utilizada nos exemplos
- [uv](https://docs.astral.sh/uv/) - Gerenciador de pacotes e ambientes Python extremamente rápido
- [jq](https://jqlang.github.io/jq/) - Processador de JSON via linha de comando

## Instalação

### 1. Instalar o uv

O [uv](https://github.com/astral-sh/uv) é um gerenciador de pacotes Python escrito em Rust, extremamente rápido e moderno.

**Linux/macOS:**
```sh
curl -LsSf https://astral.sh/uv/install.sh | sh
```

**Windows (PowerShell):**
```powershell
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

Para outras opções de instalação, consulte a [documentação oficial](https://docs.astral.sh/uv/getting-started/installation/).

### 2. Instalar as dependências do projeto

Após instalar o uv, clone o repositório e instale as dependências:

```sh
# Clone o repositório
git clone <url-do-repositorio>
cd explorando-observabilidade

# Sincroniza as dependências do projeto
uv sync
```

O comando `uv sync` irá:
- Criar automaticamente um ambiente virtual (`.venv`)
- Instalar todas as dependências especificadas no `pyproject.toml`
- Garantir que o lock file (`uv.lock`) esteja atualizado

### 3. Executar os exemplos

Para executar qualquer exemplo, você pode usar:

```sh
# Usando uv run (recomendado - executa no ambiente virtual automaticamente)
uv run python dia1/app.py

# Ou ativando o ambiente virtual manualmente
source .venv/bin/activate  # Linux/macOS
# .venv\Scripts\activate   # Windows
python dia1/app.py
```

Para mais informações sobre o uv, consulte:
- [Documentação oficial do uv](https://docs.astral.sh/uv/)
- [Guia de início rápido](https://docs.astral.sh/uv/getting-started/)
- [Gerenciamento de projetos com uv](https://docs.astral.sh/uv/guides/projects/)

## Ideias futuras

- Adicionar trace
- Trace + complexo
- Visões de métricas
- Como lidar com logs mesmo que sdk ainda não tenha ponte?
- Exportar métrica para alguma ferramenta / collector
- Exportar trace para alguma ferramenta / collector
- Exportar logs para alguma ferramenta / collector
- TelemetryGen e coletor otel
- Sampling
- Processors do coletor otel
- Instrumentção manual / automática
- Tracetest
- logfire
- rastros de banco de dados
- rastros de clientes http
- https://opentelemetry.io/blog/2025/otel-weaver/

