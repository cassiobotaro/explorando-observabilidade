package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func lançarDado(w http.ResponseWriter, r *http.Request) {
	// A rolagem de dados é feita com um número aleatório entre 1 e 6
	roll := 1 + rand.Intn(6)

	// O resultado é convertido em string,
	// adicionado uma quebra de linha
	resp := strconv.Itoa(roll) + "\n"
	// A resposta é enviada ao cliente
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
