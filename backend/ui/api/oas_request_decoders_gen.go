// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeCompleteUserRegistrationRequest(r *http.Request) (
	req *CompleteUserRegistrationReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request CompleteUserRegistrationReq
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeLoginRequest(r *http.Request) (
	req *LoginReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request LoginReq
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeSendEmailVerificationRequest(r *http.Request) (
	req *SendEmailVerificationReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request SendEmailVerificationReq
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeVerifyEmailRequest(r *http.Request) (
	req *VerifyEmailReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request VerifyEmailReq
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
