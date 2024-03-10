package main

import (
	"context"
	"log"
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
	medidor = otel.Meter("app")
)

func init() {
	// Registra o momento que a aplicação é inicializada
	inicio := time.Now()
	// Inicializa o contador observável para medir a duração da execução da tarefa
	// com descrição e unidade
	var err error
	_, err = medidor.Float64ObservableCounter(
		"app.atividade.tempo",
		metric.WithDescription("A duração desde o início do aplicativo."),
		metric.WithUnit("s"),
		// Função executada para observar a métrica
		// no momento que o SDK exporta as métricas
		metric.WithFloat64Callback(func(_ context.Context, o metric.Float64Observer) error {
			// Observe grava o valor da métrica
			o.Observe(float64(time.Since(inicio).Seconds()))
			return nil
		}),
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

	// Inicia o servidor web na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
