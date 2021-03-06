package sessions

import (
	"github.com/abhinavzspace/slumber/domain"
)

type IRevokedTokenRepositoryFactory interface {
	New(db domain.IDatabase) IRevokedTokenRepository
}
type IRevokedTokenRepository interface {
	CreateRevokedToken(token IRevokedToken) error
	DeleteExpiredTokens() error
	IsTokenRevoked(id string) bool
}
