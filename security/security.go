package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"
)

// https://oapi.dingtalk.com/robot/send?access_token=xxx
const dingTalkOAPI = "oapi.dingtalk.com"

var dingTalkURL url.URL = url.URL{
	Scheme: "https",
	Host:   dingTalkOAPI,
	Path:   "robot/send",
}

// GetDingTalkURL get DingTalk URL with accessToken & secret
// If no signature is set, the secret is set to ""
// 如果没有加签，secret 设置为 "" 即可
func GetDingTalkURL(accessToken string, secret string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	return GetDingTalkURLWithTimestamp(timestamp, accessToken, secret)
}

// GetDingTalkURLWithTimestamp get DingTalk URL with timestamp & accessToken & secret
func GetDingTalkURLWithTimestamp(timestamp string, accessToken string, secret string) (string, error) {
	dtu := dingTalkURL
	value := url.Values{}
	value.Set("access_token", accessToken)

	if secret == "" {
		dtu.RawQuery = value.Encode()
		return dtu.String(), nil
	}

	sign, err := sign(timestamp, secret)
	if err != nil {
		dtu.RawQuery = value.Encode()
		return dtu.String(), err
	}

	value.Set("timestamp", timestamp)
	value.Set("sign", sign)
	dtu.RawQuery = value.Encode()
	return dtu.String(), nil
}

func sign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
