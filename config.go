package filejanitor

import (
	"time"

	"github.com/toxyl/flo"
)

type Config struct {
	Policies []struct {
		Path            string        `yaml:"path"`
		Extension       string        `yaml:"extension"`
		KeepLast        uint          `yaml:"keep_last"`
		ScanEvery       time.Duration `yaml:"scan_every"`
		RetentionPeriod time.Duration `yaml:"retention_period"`
	} `yaml:"policies"`
}

func (c *Config) AddPolicy(path, extension string, keepLast uint, scanEvery, retentionPeriod time.Duration) {
	c.Policies = append(c.Policies, struct {
		Path            string        "yaml:\"path\""
		Extension       string        "yaml:\"extension\""
		KeepLast        uint          "yaml:\"keep_last\""
		ScanEvery       time.Duration "yaml:\"scan_every\""
		RetentionPeriod time.Duration "yaml:\"retention_period\""
	}{
		Path:            path,
		Extension:       extension,
		KeepLast:        keepLast,
		ScanEvery:       scanEvery,
		RetentionPeriod: retentionPeriod,
	})
}

func (c *Config) Save(file string) error {
	return flo.File(file).StoreYAML(c)
}

func (c *Config) Load(file string) (*Config, error) {
	return c, flo.File(file).LoadYAML(c)
}

func NewConfig() *Config {
	c := &Config{
		Policies: []struct {
			Path            string        "yaml:\"path\""
			Extension       string        "yaml:\"extension\""
			KeepLast        uint          "yaml:\"keep_last\""
			ScanEvery       time.Duration "yaml:\"scan_every\""
			RetentionPeriod time.Duration "yaml:\"retention_period\""
		}{},
	}
	return c
}

func ConfigFromFile(path string) (*Config, error) {
	return NewConfig().Load(path)
}
