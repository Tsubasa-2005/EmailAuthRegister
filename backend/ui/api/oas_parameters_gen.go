// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetAllUsersParams is parameters of GetAllUsers operation.
type GetAllUsersParams struct {
	// Page number to return.
	Page int
}

func unpackGetAllUsersParams(packed middleware.Parameters) (params GetAllUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		params.Page = packed[key].(int)
	}
	return params
}

func decodeGetAllUsersParams(args [0]string, argsEscaped bool, r *http.Request) (params GetAllUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: page.
	{
		val := int(1)
		params.Page = val
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.Page = c
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.Page)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
