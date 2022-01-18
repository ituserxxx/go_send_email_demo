package main

import (
	"github.com/go-gomail/gomail"
	"log"
)

const (
	Addr     = "smtp.qq.com"
	Port = 25
	SendUser = "jstgs@qq.com"
	SendUserAuthCode  = "rivcodvfkcxqbcdi"
)

func main() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "jstgs@qq.com" /*"发件人地址"*/, "发件人") // 发件人
	m.SetHeader("To", m.FormatAddress("xxx@vtian.top", "收件人")) // 收件人
	m.SetHeader("Subject", "主题")     // 主题
	//m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的
	m.SetBody("text/html; charset=UTF-8","我是正文") // 正文
	m.Attach("golang.jpg")  //添加附件
	d := gomail.NewDialer(Addr, Port, SendUser, SendUserAuthCode) // 发送邮件服务器、端口、发件人账号、发件人密码(授权码)
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return
	}
	log.Println("done.发送成功")
}
