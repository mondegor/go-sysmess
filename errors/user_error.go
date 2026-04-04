package errors

import (
	"github.com/mondegor/go-sysmess/errors/user"
	"github.com/mondegor/go-sysmess/errors/userfast"
)

type (
	// UserError - пользовательская ошибка с поддержкой локализации.
	UserError = userfast.ProtoError

	// UserProtoError - пользовательская ошибка с поддержкой аргументов и враппинга ошибки.
	UserProtoError = user.ProtoError
)

// NewUserError - создаёт ошибку типа kind.User с поддержкой локализации.
func NewUserError(code, message string) UserError {
	return userfast.New(code, message)
}

// NewUserProto - создаёт UserProtoError для создания ошибок
// типа kind.User с поддержкой аргументов, враппинга и локализации.
func NewUserProto(code, message string) UserProtoError {
	return user.New(code, message)
}
