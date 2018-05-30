package controllers

import (
	"github.com/astaxie/beego"
	"sort"
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"fmt"
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
	fmt.Println("1:"+signature)
	fmt.Println("2:"+timestamp)
	fmt.Println("3:"+nonce)
	fmt.Println("4:"+echostr)

	keys := []string{token,timestamp,nonce}
	sort.Strings(keys)
	sign := strings.Join(keys,"")
	if hex.EncodeToString(sha1.New().Sum([]byte(sign))) != signature{
		c.Ctx.WriteString("Bad Gate Way")
	} else {
		c.Ctx.WriteString(echostr)
	}
}
