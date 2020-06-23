package gateways

import (
	"fmt"
	"os"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

type ConfigRepository struct{}

func NewConfigRepository() *ConfigRepository {
	return &ConfigRepository{}
}

func (r *ConfigRepository) Save(c *domain.Config) error {
	if err := os.MkdirAll(".circle-env", os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(".circle-env/config")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(c.Ini()))
	if err != nil {
		return err
	}

	fmt.Println("created `.circle-env/config`.")
	return nil
}
