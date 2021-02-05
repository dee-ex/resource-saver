package main

import (
  "os"
  "log"
  "net/http"

  "github.com/rs/cors"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/dee-ex/resource-saver/modules/resources"
)

func serveHTTP() {
  r := mux.NewRouter()

  // generate code route
  r.HandleFunc("/resources/",
                resources.ResourceCreationHandler).Methods("POST")

  // get link from code route
  r.HandleFunc("/resources/{code:[a-zA-Z0-9]{6,}}",
                resources.ResourceRetrievalHandler).Methods("GET")

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{http.MethodGet,
                             http.MethodPost,
                             http.MethodPut,
                             http.MethodDelete,
                             http.MethodOptions,
                            },
  })

  enhanced_r := handlers.LoggingHandler(os.Stdout, c.Handler(r))

  log.Fatal(http.ListenAndServe(":6969", enhanced_r))
}

func main() {
  serveHTTP()
}
