package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"realtor.ai/types"
)

var secretKey = "secret123"

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Panic("Error in signing token", signedToken)
		return "", err
	}
	return signedToken, nil
}

func ParseToken(signedToken string) string {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Panic("Error in token parsing", err)
		return "Failed"
	}

	if !token.Valid {
		log.Panic("Invalid Token")
		return "Failed"
	}
	return "Success"
}

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	var NewUser types.SignUpType
	json.NewDecoder(r.Body).Decode(&NewUser)
	//check if any field is empty
	if NewUser.FirstName == "" || NewUser.LastName == "" || NewUser.Email == "" || NewUser.Password == "" || NewUser.ConfirmPassword == "" {
		w.Write([]byte("Fields cannot be left empty!"))
		return
	}

	//check for the minimum length of password
	if len(NewUser.Password) < 6 {
		w.Write([]byte("Password should be atleast 6 characters!"))
		return
	}

	//check if the password and confirm password matches
	if NewUser.Password != NewUser.ConfirmPassword {
		w.Write([]byte("Passwords do not match!"))
		return
	}

	//database logic to push the data into the database

	w.Write([]byte("User Successfully Signed Up."))

}
