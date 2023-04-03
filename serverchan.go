// Package serverchan make it easier for you to use ServerChan service
package serverchan

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ServerChan define a class to serve ServerChan service
type ServerChan struct {
	gw string
}

// NewServerChan creates a new ServerChan pointer instance.
//
//	Parameters:
//	 secretKey: your ServerChan secret key, type string, required
//
//	Returns:
//	 new ServerChan pointer instance
func NewServerChan(secretKey string) *ServerChan {
	return &ServerChan{
		gw: "https://sc.ftqq.com/" + secretKey + ".send",
	}
}

// Send message to ServerChan service server.
//
//	Parameters:
//	 title: message's title, type string, required, must be escaped and content max length: 256 byte
//	 content: message's content, type string, max length: 64 kB, suport Markdown syntax
//
//	Returns:
//	 msg: message returnd from ServerChan service server
//	 err: error message
func (sc *ServerChan) Send(title, content string) (msg string, err error) {
	uVlues, err := url.ParseQuery("text=" + url.QueryEscape(title) + "&desp=" + content)
	if err != nil {
		return "error", err
	}
	resp, err := http.PostForm(sc.gw, uVlues)
	if err != nil {
		return "error", err
	}
	defer resp.Body.Close()
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error", err
	}
	t := make(map[string]interface{}, 0)
	json.Unmarshal(p, &t)
	if int(t["data"].(map[string]interface{})["errno"].(float64)) == 1024 {
		return t["errmsg"].(string), errors.New("messages of the same content can only be sent once a minute")
	}
	return "success", nil
}
