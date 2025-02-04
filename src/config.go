package src

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	OutputPath string   `yaml:"output_path"`
	Includes   []string `yaml:"included_paths"`
	Excludes   []string `yaml:"excluded_paths"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
