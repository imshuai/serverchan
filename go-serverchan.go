package serverchan

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ServerChan struct {
	gw string
}

func NewServerChan(secretKey string) *ServerChan {
	return &ServerChan{
		gw: "https://sc.ftqq.com/" + secretKey + ".send",
	}
}

func (sc *ServerChan) Send(title, content string) (string, error) {
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
	if int(t["errno"].(float64)) == 1024 {
		return t["errmsg"].(string), errors.New("messages of the same content can only be sent once a minute")
	}
	return "success", nil
}
