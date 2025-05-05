package ports

import "context"

type DBSecretProvider interface {
	ServerUrl(ctx context.Context) (string, error)
	DBName(ctx context.Context) (string, error)
}
