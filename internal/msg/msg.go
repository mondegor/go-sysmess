package msg

import "github.com/mondegor/go-sysmess/mrmsg"

type (
	// MessageFormatter - cм. mrmsg.MessageFormatter.
	MessageFormatter = mrmsg.MessageFormatter
)

// NewMessageFormatter - cм. mrmsg.NewMessageFormatter.
func NewMessageFormatter(leftDelim, rightDelim string, formatter func(placeholder string, index int) (newPlaceholder string)) *MessageFormatter {
	return mrmsg.NewMessageFormatter(leftDelim, rightDelim, formatter)
}
