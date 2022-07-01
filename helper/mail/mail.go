package mail

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type MailHelper struct {
	Service *gmail.Service
}

func NewMailHelper(at, rt, client_id, client_secret, redirect_url string) (*MailHelper, error) {
	config := oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirect_url,
	}

	token := oauth2.Token{
		AccessToken:  at,
		TokenType:    "Bearer",
		RefreshToken: rt,
		Expiry:       time.Now(),
	}

	tokenSource := config.TokenSource(context.Background(), &token)

	svc, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, err
	}

	return &MailHelper{
		Service: svc,
	}, nil
}

func (m *MailHelper) SendVerificationCode(code, recipient, verify_type string) error {
	var message gmail.Message

	var purpose string
	if verify_type == "EMAIL_VERIFY" {
		purpose = "Email Verification"
	} else if verify_type == "PASS_RESET" {
		purpose = "Reset Password Request"
	}

	emailTo := fmt.Sprintf("To: %s\r\n", recipient)
	subject := "Subject: Nomizo - " + purpose + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf("Please input this code within 5 minute: <br><br> <b>%s</b>", code)

	msg := []byte(emailTo + subject + mime + "\n" + body)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	_, err := m.Service.Users.Messages.Send("me", &message).Do()
	return err
}
