package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdk "go.opentelemetry.io/otel/sdk/metric"
)

var (
	// Meter retorna um novo medidor com o nome especificado do fornecedor de métricas global.
	// Se um medidor com o mesmo nome já existir, o mesmo medidor será retornado.
	// O nome deve ser o nome da biblioteca ou aplicação que está sendo instrumentada.
	medidor = otel.Meter("tarefa")

	// Define um histograma para medir a duração da execução da tarefa
	histograma metric.Float64Histogram
)

func init() {
	// Inicializa o histograma para medir a duração da execução da tarefa
	// com descrição e unidade
	var err error
	histograma, err = medidor.Float64Histogram(
		"tarefa.duracao",
		metric.WithDescription("A duração da execução da tarefa."),
		metric.WithUnit("s"),
	)
	// A métrica DEVE ser registrada, senão o programa irá falhar
	if err != nil {
		panic(err)
	}
}

func main() {
	// Cria um novo exportador de métricas para stdout (console, no nosso caso)
	// A opção WithPrettyPrint() faz com que as métricas sejam impressas
	// de forma mais legível
	exportadorMetricas, err := stdoutmetric.New(
		stdoutmetric.WithPrettyPrint(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Cria um novo fornecedor de métricas que é responsável por gerenciar
	//os recursos necessários para coletar e exportar métricas
	fornecedorMedidores := sdk.NewMeterProvider(
		// PeriodicReader é um leitor de métricas que coleta métricas
		// em intervalos regulares e as exporta para o fornecedor de métricas
		sdk.WithReader(sdk.NewPeriodicReader(exportadorMetricas,
			// Para fins demonstrativos, alteramos o intervalo de coleta
			// de métricas para 10 segundos
			// O padrão é 1 minuto
			sdk.WithInterval(10*time.Second))),
	)
	// Registra o fornecedor de métricas como global
	// Se um MeterProvider não for criado, as APIs do OpenTelemetry para métricas usarão uma implementação no-op e não conseguirão gerar dados.
	otel.SetMeterProvider(fornecedorMedidores)

	// Registra um manipulador para a rota raiz do servidor web que simula a execução de uma tarefa
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Registra o tempo de início da tarefa
		inicio := time.Now()

		// Simula uma tarefa que leva um tempo aleatório para ser concluída
		duraçãoTarefa := time.Duration(rand.Intn(10)+1) * time.Second
		time.Sleep(duraçãoTarefa)

		// Registra o tempo de término da tarefa
		duration := time.Since(inicio)

		// O registro agrega um valor adicional à distribuição.
		histograma.Record(r.Context(), duration.Seconds())
	})
	// Inicia o servidor web na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
