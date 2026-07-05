package errors

import (
	"github.com/mondegor/go-sysmess/errors/user"
	"github.com/mondegor/go-sysmess/errors/userfast"
)

type (
	// UserError - прототип пользовательской ошибки с поддержкой обёртывания и локализации.
	// Используется для пользовательских ошибок без аргументов.
	UserError = userfast.ProtoError

	// UserProtoError - прототип пользовательской ошибки с поддержкой аргументов, обёртывания и локализации.
	UserProtoError = user.ProtoError
)

// NewUserError - создаёт пользовательскую ошибку с кодом и сообщением.
func NewUserError(code, message string) UserError {
	return userfast.New(code, message)
}

// NewUserProto - создаёт прототип пользовательской ошибки с поддержкой аргументов.
func NewUserProto(code, message string) UserProtoError {
	return user.New(code, message)
}
