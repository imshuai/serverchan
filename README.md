# Server Chan golang SDK

[![GoDoc](https://godoc.org/github.com/imshuai/serverchan?status.svg)](https://godoc.org/github.com/imshuai/serverchan)
## 安装

`go get -v -u github.com/imshuai/serverchan`

## 注意

1.  消息标题`title`必须经过转义，内容大小为256 Byte，必须参数
2.  消息内容`content`不需要转义，内容大小为64 KByte，支持`Markdown`语法，可选，默认为空
3.  同样内容的消息一分钟只能发送成功一次

## 使用方法
1. 访问 [http://sc.ftqq.com](http://sc.ftqq.com), 使用你的Github账号登录
2. 访问 [http://sc.ftqq.com/?c=wechat&a=bind](http://sc.ftqq.com/?c=wechat&a=bind), 使用微信扫描二维码绑定微信账号以接收消息
3. 访问 [http://sc.ftqq.com/?c=code](http://sc.ftqq.com/?c=code) 获取你的ServerChan密匙`secret key`
4. 按照下面的示例代码编写你自己的消息通知程序

## 示例代码
```c#
package main

import (
	"log"

	"github.com/imshuai/serverchan"
)

const (
	secretKey = "your serverchan secret key"
)

func main() {
	sc := serverchan.NewServerChan(secretKey)
	msgTitle := "message title with urlescape"
	msgContent := "message content"
	msgReturn, err := sc.Send(msgTitle, msgContent)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message:", msgReturn)
}
```
## license

[Apache License version 2.0](http://www.apache.org/licenses/)
