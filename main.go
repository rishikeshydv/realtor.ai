package realtorai

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"realtor.ai/auth"
	"realtor.ai/health"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", health.HealthCheck).Methods("GET") //health check

	r.HandleFunc("/api/v1/signup", auth.SignUpUser).Methods("POST")    //signup route
	r.HandleFunc("/api/v1/login", auth.LoginHandler).Methods("POST")   //login route
	r.HandleFunc("/api/v1/logout", auth.Logout).Methods("GET")         //logout route
	r.HandleFunc("/api/v1/protected", auth.CheckCookie).Methods("GET") //protected route

	log.Println("Server Running on port 5002")
	log.Fatal(http.ListenAndServe(":5002", r))

}
