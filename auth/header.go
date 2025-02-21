package auth

import (
	"net/http"
)

//this function checks if the token is valid or not
//this helps to confirm if the user is still logged in

func CheckCookie(w http.ResponseWriter, r *http.Request) {
	//get the cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		w.Write([]byte("User not logged in."))
		return
	}
	//get the token
	tokenString := cookie.Value
	returnText := ParseToken(tokenString)
	if returnText == "Failed" {
		w.Write([]byte("User Not Logged In"))
	} else {
		w.Write([]byte("User Logged In"))
	}
}
