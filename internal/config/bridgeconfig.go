package config

import "maunium.net/go/mautrix/bridge/bridgeconfig"

// implement mautrix/config BridgeConfig interface
type BridgeConfig struct {
	MessageStatusEvents bool `yaml:"message_status_events"`
	MessageErrorNotices bool `yaml:"message_error_notices"`

	CommandPrefix string `yaml:"command_prefix"`

	ManagementRoomText bridgeconfig.ManagementRoomTexts `yaml:"management_room_text"`
	Encryption         bridgeconfig.EncryptionConfig    `yaml:"encryption"`
	ResendBridgeInfo   bool                             `yaml:"resend_bridge_info"`
}

func (bc BridgeConfig) GetEncryptionConfig() bridgeconfig.EncryptionConfig {
	return bc.Encryption
}

func (bc BridgeConfig) EnableMessageStatusEvents() bool {
	return bc.MessageStatusEvents
}

func (bc BridgeConfig) EnableMessageErrorNotices() bool {
	return bc.MessageErrorNotices
}

func (bc BridgeConfig) GetCommandPrefix() string {
	return bc.CommandPrefix
}

func (bc BridgeConfig) GetManagementRoomTexts() bridgeconfig.ManagementRoomTexts {
	return bc.ManagementRoomText
}

func (bc BridgeConfig) GetResendBridgeInfo() bool {
	return bc.ResendBridgeInfo
}
