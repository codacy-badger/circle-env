package usecases

import (
	"github.com/kou-pg-0131/circle-env/src/domain"
)

// IConfigRepository ...
type IConfigRepository interface {
	Save(c *domain.Config) error
	Get() (*domain.Config, error)
}
