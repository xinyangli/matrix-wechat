package config

import "maunium.net/go/mautrix/bridge/bridgeconfig"

type Config struct {
	*bridgeconfig.BaseConfig
	Bridge BridgeConfig `yaml:"bridge"`
}
