package types

type AppServer interface {
	Serve() error
	Close() error
	Addr() string
}
