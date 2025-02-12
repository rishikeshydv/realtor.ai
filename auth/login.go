package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"realtor.ai/types"
)

var DummyEmail = "test@email.com"
var DummyPass = "testpassword"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var UserInfo types.LoginType
	err := json.NewDecoder(r.Body).Decode(&UserInfo)
	if err != nil {
		log.Panic("Error in encoding request user info", err)
	}
	//write a database logic to check if the user exists in the database

	//for now i'll use a dummy email and password
	if UserInfo.Email == DummyEmail && UserInfo.Password == DummyPass {
		//if success, create a token
		signedToken, err := CreateToken(UserInfo.Email)
		if err != nil {
			w.Write([]byte("Error occurred while creating tokens."))
		}
		w.Header().Set("Authorization", "Bearer "+signedToken)
		w.Write([]byte("User Logged In."))
	}

	w.Write([]byte("Error Logging In."))

}
