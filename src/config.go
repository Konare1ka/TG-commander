package main

import (
	"os"
	"path/filepath"
	"gopkg.in/yaml.v2"
)

var cfg *Config

type Config struct {
	Token string `yaml:"token"`
	Users []string `yaml:"users"`
	AllowsPlugins []string `yaml:"allows"`
	PluginPath string `yaml:"pluginsPath"`
}

func config() {
	execPath, err := os.Executable()
	if err != nil {logger.Error("Failed get a execPath", "err", err.Error()); os.Exit(1)}
	configPath := filepath.Join(filepath.Dir(execPath), "config.yaml")
	logger.Debug("Config", "path", configPath)

	_, err = os.Stat(configPath)
	if err != nil {
		logger.Warn("Config file not exists")
		makeConfig(configPath) 
		os.Exit(1)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {logger.Error("Failed to read data from config file", "err", err.Error()); os.Exit(1)}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {logger.Error("Failed parse YAML", "err", err.Error()); os.Exit(1)}
	logger.Debug("Config struct have", "cfg", cfg)
	logger.Info("Successful get config")

}

func makeConfig(configPath string) {
	content := `token: "your_token"
users: ["your_user"]
allows: ["start"]
pluginsPath: ""`

	err := os.WriteFile(configPath, []byte(content), 0644)
	if err != nil { logger.Error("Failed to create config.yaml", "err", err.Error()) }
}