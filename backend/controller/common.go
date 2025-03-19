package controller

import (
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
//? sending email functionality
func SendOTPWithEmail(HTML, Reciver,  string) bool {
	apiKey := os.Getenv("RESEND_API_KEY")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{"delivered@resend.dev"},
		Html:    HTML,
		Subject: "COnfirm your flight",
		Cc:      []string{"amaanmirza282@gmail.com"},
		Bcc:     []string{"amaanmirza287@gmail.com"},
		ReplyTo: "amaanmirza287@gmail.com",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(sent.Id)
	return true
}
*/
