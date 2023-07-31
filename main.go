package main

import (
	"flag"
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

var homeserver = flag.String("homeserver", "", "Matrix homeserver")
var username = flag.String("username", "", "Matrix username localpart")
var password = flag.String("password", "", "Matrix password")
var database = flag.String("database", "mautrix-example.db", "SQLite database path")
var debug = flag.Bool("debug", false, "Enable debug logs")

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		sender, _ := msg.Sender()
		fmt.Println(sender.AvatarID(), msg, msg.Content)
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()

	// 执行热登录
	bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
