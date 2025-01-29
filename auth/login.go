package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"realtor.ai/types"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var UserInfo types.UserType
	err := json.NewDecoder(r.Body).Decode(&UserInfo)
	if err != nil {
		log.Panic("Error in encoding request user info", err)
	}

}
