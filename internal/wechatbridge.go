package internal

import (
	"github.com/xinyangli/matrix-wechat/internal/config"
	"maunium.net/go/mautrix/bridge"
)

type WechatBridge struct {
	bridge.Bridge
	Config        *config.Config
	ExampleConfig string
}

func (br *WechatBridge) Init() {

}

func (br *WechatBridge) GetExampleConfig() {

}
