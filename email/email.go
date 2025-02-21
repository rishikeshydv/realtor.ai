package email

//here we link gmail and LLM
//the user should get a notification that he needs to reply
//we should have our own UI for email
//fill the space with the LLM generated text
//use hugging face models in langchain to get the LLM done

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func RunMailService() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}

	//send email
	from := "your-email@gmail.com"
	to := "recipient-email@gmail.com"
	ret, err := SendMail(from, to, "Gmail With Go", "It worked!", srv)
	if err != nil {
		log.Panic("Error")
	}
	fmt.Println("Email sent:", ret)
}

// function to get user email and connect it's email to the backend
func ConnectUserEmail(user_email string) string {
	//verify the email type
	addr, err := mail.ParseAddress(user_email)
	if err != nil {
		log.Println("Invalid Email Entry")
		return ""
	}
	return addr.Address

}

// function to get all the inbox emails
func ReadInboxEmails() {

}

// function to send every new email through the LLM
func NewEmailEnquiry() {

}
