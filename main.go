package realtorai

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"realtor.ai/health"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", health.HealthCheck).Methods("GET") //health check

	log.Println("Server Running on port 5002")
	log.Fatal(http.ListenAndServe(":5002", r))

}
