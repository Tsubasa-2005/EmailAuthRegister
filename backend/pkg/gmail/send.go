package gmail

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (c *Client) SendEmail(to, subject, body string) error {
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		c.From, to, subject, body)

	if err := c.Connect.Mail(c.From); err != nil {
		return errors.Wrap(err, "failed to set sender")
	}
	if err := c.Connect.Rcpt(to); err != nil {
		return errors.Wrap(err, "failed to set recipient")
	}

	wc, err := c.Connect.Data()
	if err != nil {
		return errors.Wrap(err, "failed to open message writer")
	}
	if _, err = wc.Write([]byte(msg)); err != nil {
		return errors.Wrap(err, "failed to write message")
	}
	if err = wc.Close(); err != nil {
		return errors.Wrap(err, "failed to close message writer")
	}

	return nil
}
