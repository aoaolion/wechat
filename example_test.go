package wechat_test

import (
	"net/http"

	"github.com/esap/wechat" // 微信SDK包
	"github.com/labstack/echo"
)

func Example() {
	wechat.Debug = true
	wechat.Set("yourToken", "yourAppID", "yourSecret", "yourEncodingAesKey")
	http.HandleFunc("/", HttpHandler)
	http.ListenAndServe(":9090", nil)
}

func Example_Echo() {
	e := echo.New()
	e.Any("/", EchoHandler)
	e.Start(":9090")
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := wechat.VerifyURL(w, r)

	// 根据消息类型主动回复
	switch ctx.Msg.MsgType {
	case wechat.TypeText:
		ctx.NewText(ctx.Msg.Content).Reply() // 回复文字
	case wechat.TypeImage:
		ctx.NewImage(ctx.Msg.MediaId).Reply() // 回复图片
	case wechat.TypeVoice:
		ctx.NewVoice(ctx.Msg.MediaId).Reply() // 回复语音
	case wechat.TypeVideo:
		ctx.NewVideo(ctx.Msg.MediaId, "video title", "video description").Reply() //回复视频
	case wechat.TypeFile:
		ctx.NewFile(ctx.Msg.MediaId).Reply() // 回复文件，仅企业微信可用
	default:
		ctx.NewText("其他消息类型" + ctx.Msg.MsgType).Reply() // 回复模板消息
	}
}
func EchoHandler(c echo.Context) error {
	ctx := wechat.VerifyURL(c.Response().Writer, c.Request())

	// 根据消息类型主动回复
	switch ctx.Msg.MsgType {
	case wechat.TypeText:
		ctx.NewText(ctx.Msg.Content).Reply() // 回复文字
	case wechat.TypeImage:
		ctx.NewImage(ctx.Msg.MediaId).Reply() // 回复图片
	case wechat.TypeVoice:
		ctx.NewVoice(ctx.Msg.MediaId).Reply() // 回复语音
	case wechat.TypeVideo:
		ctx.NewVideo(ctx.Msg.MediaId, "video title", "video description").Reply() //回复视频
	case wechat.TypeFile:
		ctx.NewFile(ctx.Msg.MediaId).Reply() // 回复文件，仅企业微信可用
	default:
		ctx.NewText("其他消息类型" + ctx.Msg.MsgType).Reply() // 回复模板消息
	}
	return nil
}