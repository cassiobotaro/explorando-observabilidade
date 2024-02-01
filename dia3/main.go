package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdk "go.opentelemetry.io/otel/sdk/metric"
)

var (
	// Cria um medidor para o pacote "api"
	medidorApi = otel.Meter("api")
	// Cria um contador para o número de chamadas
	contadorChamadas metric.Int64Counter
)

func init() {
	// Inicializa o contador para o número de chamadas
	// com descrição e unidade
	var err error
	contadorChamadas, err = medidorApi.Int64Counter(
		"api.chamadas",
		metric.WithDescription("Número de chamadas"),
		// A unidade é utilizada para especificar a unidade de medida
		// da métrica
		// Utilizamos {} para indicar que a pluralização da unidade
		metric.WithUnit("{chamada}"),
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

	// Registra rota principal que contabiliza o número de chamadas
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		contadorChamadas.Add(r.Context(), 1)
	})

	// Registra rota para adicionar valores a métrica com atributos
	// Isto servirá para demonstrar como agrupar valores de métricas
	http.HandleFunc("/rand", func(w http.ResponseWriter, r *http.Request) {

		// Gera um número aleatório entre 0 e 2
		aleatorio := rand.Intn(3)

		// Acrescenta 1 ao contador de chamadas, porém com um atributo
		contadorChamadas.Add(r.Context(), 1, metric.WithAttributes(
			// O atributo é utilizado para agrupar os valores da métrica
			attribute.Int("rand", aleatorio)),
		)

		// O resultado é convertido em string,
		// adicionado uma quebra de linha
		resp := strconv.Itoa(aleatorio) + "\n"
		// A resposta é enviada ao cliente
		if _, err := io.WriteString(w, resp); err != nil {
			log.Printf("Write failed: %v\n", err)
		}

	})

	// Inicia o servidor web na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
