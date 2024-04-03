package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdk "go.opentelemetry.io/otel/sdk/metric"
)

// registrarMetricasDoBanco registra métricas assíncronas para o banco de dados fornecido.
// Certifique-se de cancelar o registro do metric.Registration antes de fechar o banco de dados fornecido.
func registrarMetricasDoBanco(db *sql.DB, meter metric.Meter) (metric.Registration, error) {

	// Cria um contador que é usado para medir o número máximo de conexões abertas
	max, err := meter.Int64ObservableUpDownCounter(
		"db.client.connections.max",
		metric.WithDescription("O número máximo de conexões abertas no banco de dados."),
		metric.WithUnit("{connection}"),
	)
	if err != nil {
		return nil, err
	}

	tempoEspera, err := meter.Int64ObservableUpDownCounter(
		"db.client.connections.wait_time",
		metric.WithDescription("O tempo em ms que levou para obter uma conexão aberta do pool."),
		metric.WithUnit("ms"),
	)
	if err != nil {
		return nil, err
	}

	// Registra um callback que coleta as métricas do banco de dados
	// Será chamado a cada intervalo de coleta de métricas
	reg, err := meter.RegisterCallback(
		func(_ context.Context, o metric.Observer) error {
			stats := db.Stats()
			// Atualiza o contador com o número máximo de conexões abertas
			o.ObserveInt64(max, int64(stats.MaxOpenConnections))
			// Atualiza o contador com o tempo de espera para obter uma conexão
			o.ObserveInt64(tempoEspera, int64(stats.WaitDuration))
			return nil
		},
		max,
		tempoEspera,
	)
	if err != nil {
		return nil, err
	}
	return reg, nil
}

func main() {

	// Abre um banco de dados SQLite em um arquivo chamado test.db
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Meter retorna um novo medidor com o nome especificado do fornecedor de métricas global.
	// Se um medidor com o mesmo nome já existir, o mesmo medidor será retornado.
	// O nome deve ser o nome da biblioteca ou aplicação que está sendo instrumentada.
	meter := otel.Meter("db.client")

	// Registra métricas para o banco de dados
	reg, err := registrarMetricasDoBanco(db, meter)
	if err != nil {
		log.Fatal(err)
	}
	// Cancela o registro das métricas do banco de dados e fecha o banco de dados
	// Esta ordem é importante para não haver coleta de métricas após o fechamento do banco de dados
	defer reg.Unregister()
	defer db.Close()

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

	// A ideia do servidor web aqui é apenas manter o programa em execução
	log.Fatal(http.ListenAndServe(":8080", nil))

}
