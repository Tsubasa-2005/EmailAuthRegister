package email

type VerificationTemplate struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
