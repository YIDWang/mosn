package jwtauthn

import "errors"

var (
	ErrInvalidToken          = errors.New("invalid token")
	ErrJwtNotFound           = errors.New("jwt token not found")
	ErrJwtBadFormat          = errors.New("jwt bad format")
	ErrJwtAudienceNotAllowed = errors.New("jwt audience not allowed")
	ErrJwtExpired            = errors.New("jwt audience not allowed")
	ErrJwtNotYetValid        = errors.New("jwt not yet valid")
	ErrJwksNoValidKeys       = errors.New("jwks no valid keys")
	ErrJwtUnknownIssuer      = errors.New("jwt unknown issuer")
	ErrJwksFetch             = errors.New("failed to fetch jwks")
)
