package mailer

type Mailer interface {
	SendEmail(*Mail) error
}

type Mail struct {
	To string `json:"to" validate:"required,email"`
	Message string `json:"message" validate:"required"`
	Subject string `json:"subject"`
}
