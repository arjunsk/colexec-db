package client

import "context"

type SQLExecutor interface {
	Exec(ctx context.Context, sql string) (Result, error)
	ExecTxn(ctx context.Context, execFunc func(TxnExecutor) error) error
}
type TxnExecutor interface {
	Exec(sql string) (Result, error)
}
