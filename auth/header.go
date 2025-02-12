package auth

import (
	"net/http"
)

//this function checks the header of each request if the 'Authorization' exists or not
//use the same function to check for the protected routes

func CheckHeader(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.Write([]byte("User not logged in."))
		return
	}
	tokenString = tokenString[len("Bearer "):]
	returnText := ParseToken(tokenString)
	if returnText == "Failed" {
		w.Write([]byte("User Not Logged In"))
	}
	w.Write([]byte("User Successfully logged in."))
}
