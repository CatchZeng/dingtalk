package dingtalk

import (
	"bou.ke/monkey"
	"errors"
	"github.com/CatchZeng/dingtalk/internal/security"
	"io"
	"net/http"
	"reflect"
	"testing"

	mock_message "github.com/CatchZeng/dingtalk/test/mocks/message"
	"github.com/golang/mock/gomock"
)

func TestNewClient(t *testing.T) {
	type args struct {
		accessToken string
		secret      string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "",
			args: args{
				accessToken: "123456",
				secret:      "111111",
			},
			want: &Client{
				AccessToken: "123456",
				Secret:      "111111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.accessToken, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	messgae := mock_message.NewMockMessage(ctrl)

	t.Run("message return error", func(t *testing.T) {
		c := &Client{}

		messgae.EXPECT().ToByte().Return([]byte{}, errors.New("test"))

		if _, err := c.Send(messgae); err == nil {
			t.Error("send error")
		}
	})

	t.Run("security.GetDingTalkURL return error", func(t *testing.T) {
		c := &Client{
			AccessToken: "dasfsafewfewfwfewf",
			Secret: "ewfewfwfwefwafew",
		}

		messgae.EXPECT().ToByte().Return([]byte{}, nil)
		monkey.Patch(security.GetDingTalkURL, func(accessToken string, secret string) (string, error) {
			return "", errors.New("GetDingTalkURL error")
		})

		if _, err := c.Send(messgae); err == nil {
			t.Error("send error")
		}
	})


	t.Run("http.NewRequest return error", func(t *testing.T) {
		c := &Client{
			AccessToken: "dasfsafewfewfwfewf",
			Secret: "ewfewfwfwefwafew",
		}

		messgae.EXPECT().ToByte().Return([]byte{}, nil)
		monkey.Patch(security.GetDingTalkURL, func(accessToken string, secret string) (string, error) {
			return "https://oapi.dingtalk.com/robot/send?access_token=ewfewfwfwefwafew", nil
		})

		monkey.Patch(http.NewRequest, func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("NewRequest error")
		})

		if _, err := c.Send(messgae); err == nil {
			t.Error("send error")
		}
	})
}
