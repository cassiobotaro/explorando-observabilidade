package main

import (
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdk "go.opentelemetry.io/otel/sdk/metric"
)

var (
	// Cria um medidor para o pacote "itens"
	medidorItens = otel.Meter("itens")
	// Cria um contador para o número de chamadas
	contadorItens metric.Int64UpDownCounter
)

func init() {
	// Inicializa o contador para o número de itens
	// com descrição e unidade
	var err error
	contadorItens, err = medidorItens.Int64UpDownCounter(
		"itens.contador",
		metric.WithDescription("Número de itens"),
		// A unidade é utilizada para especificar a unidade de medida
		// da métrica
		// Utilizamos {} para indicar que a pluralização da unidade
		metric.WithUnit("{item}"),
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

	// Registra rota de incremento de itens
	http.HandleFunc("/adicionaItem", func(w http.ResponseWriter, r *http.Request) {
		// Quando um item é adicionado, incrementamos o contador
		contadorItens.Add(r.Context(), 1)
	})

	// Registra rota de decremento de itens
	http.HandleFunc("/removeItem", func(w http.ResponseWriter, r *http.Request) {
		// Quando um item é removido, decrementamos o contador
		contadorItens.Add(r.Context(), -1)
	})

	// Inicia o servidor web na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
