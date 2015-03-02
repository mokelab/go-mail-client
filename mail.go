package go_mail_client

type MailClient interface {
	// Sends email
	Send(email, subject, body string) error
}
