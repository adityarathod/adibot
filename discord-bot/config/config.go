package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Allowlist struct {
	Enabled bool     `json:"enabled"`
	IdList  []string `json:"ids"`
	Ids     map[string]bool
}

type BotConfig struct {
	Token            string    `json:"token"`
	UserAllowlist    Allowlist `json:"userAllowlist"`
	ChannelAllowlist Allowlist `json:"channelAllowlist"`
	ReplyRatelimit   struct {
		Enabled    bool    `json:"enabled"`
		Proportion float64 `json:"proportion"`
	} `json:"replyRatelimit"`
	ModelEndpoint string `json:"modelEndpoint"`
}

func (c *BotConfig) IsUserAllowlisted(userID string) bool {
	if !c.UserAllowlist.Enabled {
		return true
	}
	return c.UserAllowlist.Ids[userID]
}

func (c *BotConfig) IsChannelAllowlisted(channelID string) bool {
	if !c.ChannelAllowlist.Enabled {
		return true
	}
	return c.ChannelAllowlist.Ids[channelID]
}

func LoadBotConfig(configPath string) (BotConfig, error) {
	config := BotConfig{}
	if configPath == "" {
		configPath = path.Join(".", "bot-config.json")
	}
	configJson, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("failed to read bot config file: %w", err)
	}
	if err := json.Unmarshal(configJson, &config); err != nil {
		return config, fmt.Errorf("failed to unmarshal bot config: %w", err)
	}
	// Convert the allowlist IDs to a map for faster lookups
	if config.UserAllowlist.Enabled {
		config.UserAllowlist.Ids = make(map[string]bool)
		for _, id := range config.UserAllowlist.IdList {
			config.UserAllowlist.Ids[id] = true
		}
	}
	if config.ChannelAllowlist.Enabled {
		config.ChannelAllowlist.Ids = make(map[string]bool)
		for _, id := range config.ChannelAllowlist.IdList {
			config.ChannelAllowlist.Ids[id] = true
		}
	}
	return config, nil
}
