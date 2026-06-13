package ai

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configPath = ".ghx/config.json"

type Config struct {
	AI AIConfig `json:"ai"`
	UI UIConfig `json:"ui"`
}

type AIConfig struct {
	ActiveProvider string                    `json:"active_provider"`
	Providers      map[string]ProviderConfig `json:"providers"`
}

type ProviderConfig struct {
	Host    string         `json:"host,omitempty"`
	Model   string         `json:"model"`
	APIKey  string         `json:"api_key,omitempty"`
	BaseURL string         `json:"base_url,omitempty"`
	Options map[string]any `json:"options,omitempty"`
}

type UIConfig struct {
	Theme             string `json:"theme"`
	ShowAISuggestions bool   `json:"show_ai_suggestions"`
}

func DefaultConfig() Config {
	return Config{
		AI: AIConfig{
			ActiveProvider: "ollama",
			Providers: map[string]ProviderConfig{
				"ollama":   {Host: "http://localhost:11434", Model: "llama3", Options: map[string]any{"temperature": 0.7, "num_ctx": 4096}},
				"openai":   {Model: "gpt-4o"},
				"claude":   {Model: "claude-3-5-sonnet-20241022"},
				"lmstudio": {Host: "http://localhost:1234", Model: "local-model"},
				"mlx":      {Host: "http://localhost:8080"},
			},
		},
		UI: UIConfig{Theme: "dark", ShowAISuggestions: true},
	}
}

func LoadConfig() (Config, error) {
	config := DefaultConfig()
	data, err := os.ReadFile(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return config, nil
		}
		return config, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return DefaultConfig(), err
	}
	return config.withDefaults(), nil
}

func SaveConfig(config Config) error {
	if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(config.withDefaults(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, append(data, '\n'), 0o600)
}

func (c Config) ActiveProvider() ProviderConfig {
	provider, ok := c.AI.Providers[c.AI.ActiveProvider]
	if !ok {
		return ProviderConfig{}
	}
	return provider
}

func (c Config) withDefaults() Config {
	defaults := DefaultConfig()
	if c.AI.ActiveProvider == "" {
		c.AI.ActiveProvider = defaults.AI.ActiveProvider
	}
	if c.AI.Providers == nil {
		c.AI.Providers = map[string]ProviderConfig{}
	}
	for name, provider := range defaults.AI.Providers {
		if _, ok := c.AI.Providers[name]; !ok {
			c.AI.Providers[name] = provider
		}
	}
	if c.UI.Theme == "" {
		c.UI.Theme = defaults.UI.Theme
	}
	return c
}
