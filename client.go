package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CatchZeng/dingtalk/internal/security"
	"io/ioutil"
	"net/http"
	"time"
)

// Client dingtalk client
type Client struct {
	AccessToken string
	Secret      string
}

// NewClient new dingtalk client
func NewClient(accessToken, secret string) *Client {
	return &Client{
		AccessToken: accessToken,
		Secret: secret,
	}
}

// Response response struct
type Response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}

const httpTimoutSecond = time.Duration(30) * time.Second

// Send message
func (d *Client) Send(message Message) (Response, error) {
	res := Response{}

	reqBytes, err := message.ToByte()
	if err != nil {
		return res, err
	}

	pushURL, err := security.GetDingTalkURL(d.AccessToken, d.Secret)
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest(http.MethodPost, pushURL, bytes.NewReader(reqBytes))
	if err != nil {
		return res, err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	client.Timeout = httpTimoutSecond
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}

	if res.ErrCode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.ErrMsg)
	}

	return res, nil
}
