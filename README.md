# Server Chan golang SDK

## install

`go get -v -u github.com/imshuai/go-serverchan`

## attention

1.  message title must be escaped and content max length: 256 byte, required
2.  message content max length: 64 kB, suport Markdown syntax, nullable
3.  messages of the same content can only be sent once a minute

## useage
1. visit [http://sc.ftqq.com](http://sc.ftqq.com), use your github account to login
2. visit [http://sc.ftqq.com/?c=wechat&a=bind](http://sc.ftqq.com/?c=wechat&a=bind), use your Wechat APP scan the QRcode and bind your WeChat account to ServerChan
3. visit [http://sc.ftqq.com/?c=code](http://sc.ftqq.com/?c=code) to your ServerChan secret key
4. follow the example bellow and send your first message to your WeChat APP

## example
```
package main

import (
	"log"

	"github.com/imshuai/go-serverchan"
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

This package is made available under an [Apache License v2](http://www.apache.org/licenses/)
