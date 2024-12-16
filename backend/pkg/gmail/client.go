package gmail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/taxio/errors"
)

type Client struct {
	Connect *smtp.Client
	From    string
}

func NewConnect(hostname string, port int, username, password string) (*Client, error) {
	if hostname == "" || port == 0 || username == "" || password == "" {
		return nil, errors.New("invalid SMTP configuration")
	}

	serverAddr := fmt.Sprintf("%s:%d", hostname, port)
	client, err := smtp.Dial(serverAddr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to server")
	}

	if err := client.StartTLS(&tls.Config{ServerName: hostname}); err != nil {
		err := client.Close()
		if err != nil {
			return nil, errors.Wrap(err, "failed to close connection")
		}
		return nil, errors.Wrap(err, "failed to start TLS")
	}

	auth := smtp.PlainAuth("", username, password, hostname)
	if err := client.Auth(auth); err != nil {
		err := client.Close()
		if err != nil {
			return nil, errors.Wrap(err, "failed to close connection")
		}
		return nil, errors.Wrap(err, "failed to authenticate")
	}

	return &Client{
		Connect: client,
		From:    username,
	}, nil
}
