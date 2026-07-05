package db_test

import (
	"context"

	"github.com/mondegor/go-sysmess/mrstorage"
)

// fakeConnManager - тестовый двойник mrstorage.DBConnManager.
// Реализует только Conn; остальные методы наследуются от вложенного nil-интерфейса
// и паникуют при неожиданном вызове.
type fakeConnManager struct {
	mrstorage.DBConnManager
	conn *fakeConn
}

func (m *fakeConnManager) Conn(context.Context) mrstorage.DBConn {
	return m.conn
}

// fakeConn - тестовый двойник mrstorage.DBConn, возвращающий заданную ошибку
// из ExecRow; перехватывает переданные аргументы.
type fakeConn struct {
	mrstorage.DBConn
	err      error
	lastArgs []any
}

func (c *fakeConn) ExecRow(_ context.Context, _ string, args ...any) error {
	c.lastArgs = args

	return c.err
}
