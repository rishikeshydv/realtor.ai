package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"realtor.ai/db"
	"realtor.ai/types"
)

func CreateToken(username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET_KEY")
	secretKey := []byte(secret)
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET_KEY")
	secretKey := []byte(secret)
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
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		w.Write([]byte("Error in decoding request user info"))
	}
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
	database := db.LocalDBConnect()
	database.AutoMigrate(types.SignUpType{})

	//check if the user already exists
	tx := database.First(&NewUser, "email = ?", NewUser.Email)
	if tx.RowsAffected > 0 {
		w.Write([]byte("User already exists!"))
		return
	}

	//if the user does not exist, create a new user
	tx = database.Table("user_info").Create(&NewUser)
	if tx.Error != nil {
		w.Write([]byte("Error in creating user!"))
		return
	} else {
		w.Write([]byte("User Successfully Signed Up."))
	}

	w.Write([]byte("User Successfully Signed Up."))

}
