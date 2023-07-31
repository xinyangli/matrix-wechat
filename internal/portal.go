package internal

import "github.com/eatmoreapple/openwechat"

type User struct {
	bot *openwechat.Bot
}

func (u *User) createClient() {
	u.bot = openwechat.DefaultBot(openwechat.Desktop)
	u.bot.MessageHandler = func(msg *openwechat.Message) {
		// Dispatch messages
		if msg.IsText() {

		} else if msg.IsPicture() {

		} else if msg.IsLocation() {

		} else if msg.IsVoice() {

		} else if msg.IsSystem() {

		} else {

		}
	}
	u.bot.UUIDCallback = func(uuid string) {
		// Send url to user using matrix bot.
	}
}
func (u *User) Login() {
}
