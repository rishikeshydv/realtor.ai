package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"realtor.ai/db"
	"realtor.ai/types"
)

var DummyEmail = "test@email.com"
var DummyPass = "testpassword"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var UserInfo types.LoginType
	var dbUser types.SignUpType
	err := json.NewDecoder(r.Body).Decode(&UserInfo)
	if err != nil {
		log.Panic("Error in encoding request user info", err)
	}
	//write a database logic to check if the user exists in the database
	database := db.LocalDBConnect()
	tx := database.Table("user_info").First(&UserInfo, "email = ?", UserInfo.Email)
	if tx.RowsAffected == 0 {
		w.Write([]byte("User does not exist."))
		return
	}

	tx.Scan(&dbUser)
	if dbUser.Password == UserInfo.Password {
		tokenString, err := CreateToken(UserInfo.Email)
		if err != nil {
			log.Panic("Error in creating token", err)
			return
		}
		cookie := http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  time.Now().Add(1 * time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		w.Write([]byte("User Successfully Logged In."))
		return
	} else {
		w.Write([]byte("Error Logging In."))
		return
	}

}
