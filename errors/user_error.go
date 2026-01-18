package errors

import (
	"github.com/mondegor/go-sysmess/errors/user"
)

type (
	// UserProtoError - пользовательская ошибка.
	UserProtoError = user.ProtoError
)

// NewUserProto - создаёт UserProtoError для создания ошибок типа kind.User.
func NewUserProto(code, message string) UserProtoError {
	return user.New(code, message)
}
