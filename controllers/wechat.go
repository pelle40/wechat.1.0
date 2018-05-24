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

func (c *WechatController)Entry(){
	token := "top_learn_token"
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")

	keys := []string{token,timestamp,nonce}
	sort.Strings(keys)
	sign := strings.Join(keys,"")
	if hex.EncodeToString(sha1.New().Sum([]byte(sign))) != signature{
		c.Ctx.WriteString("bad gate way")
	} else {
		c.Ctx.WriteString(echostr)
	}
}