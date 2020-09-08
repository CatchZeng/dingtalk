package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"math"
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

// URL get DingTalk URL with accessToken & secret
// If no signature is set, the secret is set to ""
// 如果没有加签，secret 设置为 "" 即可
func URL(accessToken string, secret string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	return URLWithTimestamp(timestamp, accessToken, secret)
}

// URLWithTimestamp get DingTalk URL with timestamp & accessToken & secret
func URLWithTimestamp(timestamp string, accessToken string, secret string) (string, error) {
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

// Validate validate
// https://ding-doc.dingtalk.com/doc#/serverapi2/elzz1p
func Validate(signStr, timestamp, secret string) (bool, error) {
	t, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return false, err
	}

	timeGap := time.Since(time.Unix(t, 0))
	if math.Abs(timeGap.Hours()) > 1 {
		return false, fmt.Errorf("specified timestamp is expired")
	}

	ourSign, err := sign(timestamp, secret)
	if err != nil {
		return false, err
	}
	return ourSign == signStr, nil
}

func sign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
