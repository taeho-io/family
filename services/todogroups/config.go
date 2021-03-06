package todogroups

import (
	"github.com/taeho-io/family/services/base"
)

type Config interface {
	base.Config

	Settings() Settings
}

type DefaultConfig struct {
	base.Config

	settings Settings
}

func NewConfig(settings Settings) (cfg Config) {
	return &DefaultConfig{
		Config:   base.NewConfig(serviceName),
		settings: settings,
	}

}

func NewMockConfig() (cfg Config) {
	return &DefaultConfig{
		Config:   base.NewConfig(serviceName),
		settings: NewMockSettings(),
	}
}

func (t *DefaultConfig) Settings() Settings {
	return t.settings
}
