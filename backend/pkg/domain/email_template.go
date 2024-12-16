package domain

import (
	"bytes"
	"text/template"

	"github.com/taxio/errors"
)

type EmailVerificationData struct {
	URL string
}

func GenerateEmailVerificationBody(templateString string, data EmailVerificationData) (string, error) {
	tmpl, err := template.New("email_verification_vody").Parse(templateString)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse template string")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", errors.Wrap(err, "failed to execute template")
	}

	return buf.String(), nil
}
