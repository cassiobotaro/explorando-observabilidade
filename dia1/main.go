package main

import (
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

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
	fornecedorMedidores := metric.NewMeterProvider(
		// PeriodicReader é um leitor de métricas que coleta métricas
		// em intervalos regulares e as exporta para o fornecedor de métricas
		metric.WithReader(metric.NewPeriodicReader(exportadorMetricas,
			// Para fins demonstrativos, alteramos o intervalo de coleta
			// de métricas para 10 segundos
			// O padrão é 1 minuto
			metric.WithInterval(10*time.Second))),
	)
	// Registra o fornecedor de métricas como global
	otel.SetMeterProvider(fornecedorMedidores)

	// Cria um novo roteador para o servidor web
	rotas := http.NewServeMux()
	// Registra a rota /lançardados
	rotas.Handle(
		"/lançardados",
		// WithRouteTag anota uma métrica(adiciona atributo) com o nome da rota (http.route)
		otelhttp.WithRouteTag("/lançardados", http.HandlerFunc(lançarDados)))

	// O controlador é um middleware que coleta métricas para cada rota registrada
	// e as exporta para o provedor de métricas
	controlador := otelhttp.NewHandler(rotas, "/")

	// Inicia o servidor web na porta 8080
	log.Fatal(http.ListenAndServe(":8080", controlador))
}
