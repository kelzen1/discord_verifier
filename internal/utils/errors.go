package utils

import "errors"

var (
	ErrInteractionNotFound = errors.New("interaction_not_found")
	//ErrFoundTooMany        = errors.New("found_too_many")
	ErrStructMismatch = errors.New("struct_mismatch")
	ErrRoleNotFound   = errors.New("role_not_found")
)
