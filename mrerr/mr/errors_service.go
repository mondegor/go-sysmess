package mr

import "github.com/mondegor/go-sysmess/mrerr"

// ErrServiceOperationFailed - service operation is failed (ErrServiceOperationFailed оборачивает в эту ошибку все нераспознанные ошибки).
var ErrServiceOperationFailed = mrerr.NewKindInternal("service operation is failed")
