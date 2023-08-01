package internal

import (
	"github.com/xinyangli/matrix-wechat/internal/config"
	"maunium.net/go/mautrix/bridge"
	"maunium.net/go/mautrix/bridge/commands"
)

type WechatBridge struct {
	bridge.Bridge
	Config        *config.Config
	ExampleConfig string
}

func (br *WechatBridge) Init() {
	br.CommandProcessor = commands.NewProcessor(&br.Bridge)
	br.RegisterCommands()

}

func (br *WechatBridge) GetExampleConfig() string {
	return br.ExampleConfig
}
