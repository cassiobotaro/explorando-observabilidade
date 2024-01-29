package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter            = otel.Meter("lançardados")
	contadorRolagens metric.Int64Counter
)

func init() {
	// A função init é executada no momento da inicialização do pacote
	var err error
// contadorRolagens é um contador de métricas que conta o número de rolagens
// de dados por valor de rolagem
	contadorRolagens, err = meter.Int64Counter("dados.rolagens",
		metric.WithDescription("O número de rolagens por valor de rolagem"),
		// A unidade de medida é definida como "{lançamento}" para que o valor
		// pluralizado seja exibido corretamente
		metric.WithUnit("{lançamento}"))
	if err != nil {
		panic(err)
	}
}

func lançarDados(w http.ResponseWriter, r *http.Request) {
	// A rolagem de dados é feita com um número aleatório entre 1 e 6
	rolagem := 1 + rand.Intn(6)

	// Um atributo personalizado é criado com o valor da rolagem
	atributoValorRolagem := attribute.Int("dados.valor", rolagem)
	// O atributo personalizado "dados.valor" é usado para agrupar as métricas
	// por valor de rolagem

	// O contador de rolagens é incrementado com o atributo personalizado
	contadorRolagens.Add(r.Context(), 1, metric.WithAttributes(atributoValorRolagem))

	// O resultado é convertido em string,
	// adicionado uma quebra de linha
	resp := strconv.Itoa(rolagem) + "\n"
	// A resposta é enviada ao cliente
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
