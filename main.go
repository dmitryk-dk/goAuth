package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	appHandlers "github.com/dmitryk-dk/goAuth/handlers"
)

func main() {
	// listening port
	const port = "3000"
	// init mux router
	router := mux.NewRouter()
	// init index.html
	router.Handle("/", http.FileServer(http.Dir("./views/")))

	// response static data
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))
	fmt.Printf("Running server on port: %s\n Type Ctr-c to shutdown server.\n", port)
	portConn := fmt.Sprintf(":%s", port)
	// routes handlers
	router.Handle("/status", appHandlers.StatusHandler()).Methods("GET")
	router.Handle("/products", appHandlers.ProductsHandler()).Methods("GET")
	router.Handle("/products/{slug}/feedback", appHandlers.AddFeedbackHandler()).Methods("Post")
	router.Handle("/login", appHandlers.LoginHandler()).Methods("POST")

	prepareShutdown()
	//start server
	http.ListenAndServe(portConn, handlers.LoggingHandler(os.Stdout, router))
}

func prepareShutdown() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Got signal %d", <-sig)
		os.Exit(0)
	}()
}
