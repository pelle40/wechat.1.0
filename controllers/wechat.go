package controllers

import (
	"github.com/astaxie/beego"
	"sort"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

type WechatController struct{
	beego.Controller
}

func (c *WechatController)Entry() {
	account := c.GetString("account")
	token := beego.AppConfig.String(account+"::token")
	if token==""{
		c.Ctx.WriteString("Bad Gate Way")
		c.StopRun()
	}
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")

	keys := []string{token,timestamp,nonce}
	sort.Strings(keys)
	sign := strings.Join(keys,"")
	strsha := sha1.New()
	strsha.Write([]byte(sign))
	if hex.EncodeToString(strsha.Sum(nil)) != signature{
		c.Ctx.WriteString("Bad Gate Way")
	} else {
		c.Ctx.WriteString(echostr)
	}
}
