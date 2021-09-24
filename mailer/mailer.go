package mailer

type Mailer interface {
	SendEmail(to string, email string) error
}

type Mail struct {
	To string `json:"to"`
	Message string `json:"message"`
	Subject string `json:"subject"`
}

