package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ConfigLoader struct {
	configDir string
}

func NewConfigLoader(dir string) *ConfigLoader {
	return &ConfigLoader{configDir: dir}
}

func (c *ConfigLoader) LoadConfig(filename string) (map[string]interface{}, error) {
	path := filepath.Join(c.configDir, filename)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	yaml.Unmarshal(data, &config)
	return config, nil
}

func (c *ConfigLoader) SaveConfig(filename string, config map[string]interface{}) error {
	path := filepath.Join(c.configDir, filename)
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	ioutil.WriteFile(path, data, 0644)
	return nil
}

func (c *ConfigLoader) LoadUserConfig(username, configFile string) (map[string]interface{}, error) {
	userConfigPath := filepath.Join(c.configDir, username, configFile)
	data, err := os.ReadFile(userConfigPath)
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	yaml.Unmarshal(data, &config)
	return config, nil
}

func (c *ConfigLoader) ImportConfig(importPath string) error {
	data, err := os.ReadFile(importPath)
	if err != nil {
		return err
	}

	var config map[string]interface{}
	yaml.Unmarshal(data, &config)

	configName := filepath.Base(importPath)
	destPath := filepath.Join(c.configDir, configName)
	os.WriteFile(destPath, data, 0644)
	return nil
}
