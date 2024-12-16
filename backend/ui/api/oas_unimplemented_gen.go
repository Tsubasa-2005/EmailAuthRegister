// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CompleteUserRegistration implements CompleteUserRegistration operation.
//
// Complete user registration.
//
// POST /complete-registration
func (UnimplementedHandler) CompleteUserRegistration(ctx context.Context, req *CompleteUserRegistrationReq) (r CompleteUserRegistrationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAllUsers implements GetAllUsers operation.
//
// Get all users.
//
// GET /users
func (UnimplementedHandler) GetAllUsers(ctx context.Context) (r []User, _ error) {
	return r, ht.ErrNotImplemented
}

// Ping implements Ping operation.
//
// Check if the server is running.
//
// GET /ping
func (UnimplementedHandler) Ping(ctx context.Context) (r *PingOK, _ error) {
	return r, ht.ErrNotImplemented
}

// SendEmailVerification implements SendEmailVerification operation.
//
// Send email verification.
//
// POST /send-verification
func (UnimplementedHandler) SendEmailVerification(ctx context.Context, req *SendEmailVerificationReq) (r SendEmailVerificationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// VerifyEmail implements VerifyEmail operation.
//
// Verify email.
//
// POST /verify-email
func (UnimplementedHandler) VerifyEmail(ctx context.Context, req *VerifyEmailReq) (r VerifyEmailRes, _ error) {
	return r, ht.ErrNotImplemented
}
