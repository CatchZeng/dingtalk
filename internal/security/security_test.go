package security

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
)

const (
	timestamp   = "1582163555000"
	accessToken = "1c53e149ba5de6597ca2442f0e901fd86156780b8ac141e4a75afdc44c85ca4f"
	secret      = "SECb90923e19e58b466481e9e7b7a5b4f108a4531abde590ad3967fb29f0eae5c68"
	signed      = "BQKsG%2BQOCl%2BbYJOLc6pxDHxjVquzlZPWgvRzeN2J5zY%3D"
)

func TestURL(t *testing.T) {
	monkey.Patch(strconv.FormatInt, func(i int64, base int) string {
		return timestamp
	})

	defer monkey.Unpatch(strconv.FormatInt)

	got, err := URL(accessToken, secret)
	if err != nil {
		t.Errorf("URL() error = %v", err)
	}

	want := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v&sign=%v&timestamp=%v", accessToken, signed, timestamp)
	if got != want {
		t.Errorf("URL() = %v, want %v", got, want)
	}
}

func TestURLWithTimestamp(t *testing.T) {
	type args struct {
		accessToken string
		secret      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "without sign",
			args: args{
				accessToken: accessToken,
			},
			want:    fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v", accessToken),
			wantErr: false,
		},
		{
			name: "with sign",
			args: args{
				accessToken: accessToken,
				secret:      secret,
			},
			want:    fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v&sign=%v&timestamp=%v", accessToken, signed, timestamp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := URLWithTimestamp(timestamp, tt.args.accessToken, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLWithTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLWithTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	validateTimestamp := strconv.FormatInt(time.Now().Add(60*time.Second).Unix(), 10)

	result, err := sign(validateTimestamp, secret)
	if err != nil {
		t.Error(err)
	}

	_, err = Validate(result, strconv.FormatInt(time.Now().Add(-3601*time.Second).Unix(), 10), secret)
	if err == nil {
		t.Error("this should be err, but not")
	}

	_, err = Validate(result, strconv.FormatInt(time.Now().Add(3601*time.Second).Unix(), 10), secret)
	if err == nil {
		t.Error("this should be err, but not")
	}

	b, err := Validate(result, validateTimestamp, secret)
	if err != nil {
		t.Error(err)
	} else {
		if !b {
			t.Error("token is not the same")
		}
	}
}
