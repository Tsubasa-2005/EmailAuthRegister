package domain

import (
	"encoding/json"
	"os"

	"github.com/Tsubasa-2005/EmailAuthResister/templates/email"
	"github.com/taxio/errors"
)

const emailVerificationPath = "templates/email/verification.json"

func LoadEmailVerificationTemplates() (*email.VerificationTemplate, error) {
	file, err := os.Open(emailVerificationPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open email templates file")
	}
	defer file.Close()

	var templates email.VerificationTemplate
	if err := json.NewDecoder(file).Decode(&templates); err != nil {
		return nil, errors.Wrap(err, "failed to decode email templates")
	}

	return &templates, nil
}
