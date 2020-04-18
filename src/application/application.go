package application

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/james077/bookstore_items-api/src/clients/elasticsearch"

	"github.com/james077/bookstore_utils-go/logger"
)

var (
	router = mux.NewRouter()
)
// StartApplication da inicio al servidor de la aplicación
func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Addr: ":8081",
		// Buena práctica para establecer tiempos de espera para evitar ataques de Slowloris.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	logger.Info("Iniciando aplicación...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
