package internal

import "maunium.net/go/mautrix/bridge/commands"

func (br *WechatBridge) RegisterCommands() {
	proc := br.CommandProcessor.(*commands.Processor)
	proc.AddHandlers()
}
