package middleware

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/pkg/app"
	"github.com/go-tour/blog_service/pkg/email"
	"github.com/go-tour/blog_service/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {

	emailDeliver := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().ErrorF(c, "panic recover err: %v,", err)
				//send email
				go func() {

					err = emailDeliver.SendMail(global.EmailSetting.To, fmt.Sprintf("异常抛出,发生时间: %d", time.Now().Unix()),
						fmt.Sprintf("错误信息: %v", err))
					if err != nil {
						global.Logger.PanicF("mail.sendMail err: %v", err)
					}
				}()
				go func() {
					//ding
					cli := dingtalk.InitDingTalk(global.DingTalkSetting.To, ".")
					// 发个text类型消息
					textContent := "Tag:gin-recover-goroutine email & ding send. todo: 深入浅出embedding"
					err2 := cli.SendTextMessage(textContent)
					if err2 != nil {
						global.Logger.PanicF("ding.SendTextMessage err: %v", err)
					}
				}()
				app.NewResp(c).ErrResp(errcode.ServerErr)
				c.Abort()
			}
		}()
		c.Next()
	}
}
