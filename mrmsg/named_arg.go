package mrmsg

type (
	// NamedArg - именованный аргумент используемый в сообщениях.
	NamedArg struct {
		Name  string
		Value any
	}
)

// ValueString - преобразовывает значение аргумента в строку.
func (a *NamedArg) ValueString() string {
	return ToString(a.Value)
}
