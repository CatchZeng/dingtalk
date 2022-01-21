package dingtalk

import (
	"errors"
	"log"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/CatchZeng/dingtalk/configs"
)

func Test_newClient(t *testing.T) {
	fakeExit := func(int) {
		log.Print("fake exit")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	t.Run("getAccessToken return empty", func(t *testing.T) {
		accessToken = ""
		_, err := newClient()
		if err == nil {
			t.Error("newClient() error")
		}
	})

	t.Run("getAccessToken return token", func(t *testing.T) {
		accessToken = "123"
		client, err := newClient()
		if err != nil || client == nil {
			t.Error("newClient() error")
		}
	})
}

func Test_getAccessToken(t *testing.T) {
	t.Run("get from accessToken", func(t *testing.T) {
		accessToken = "123"
		got := getAccessToken()
		if got != accessToken {
			t.Errorf("getAccessToken() = %v, want %v", got, accessToken)
		}
	})

	t.Run("GetConfig error", func(t *testing.T) {
		accessToken = ""

		monkey.Patch(configs.GetConfig, func(key string) (string, error) {
			return "", errors.New("GetConfig error")
		})
		defer monkey.Unpatch(configs.GetConfig)

		got := getAccessToken()
		if got != "" {
			t.Errorf("getAccessToken() = %v, want %v", got, "")
		}
	})

	t.Run("get from config", func(t *testing.T) {
		accessToken = ""

		want := "123"
		monkey.Patch(configs.GetConfig, func(key string) (string, error) {
			return want, nil
		})
		defer monkey.Unpatch(configs.GetConfig)

		got := getAccessToken()
		if got != want {
			t.Errorf("getAccessToken() = %v, want %v", got, want)
		}
	})
}

func Test_getSecret(t *testing.T) {
	t.Run("get from secret", func(t *testing.T) {
		secret = "123"
		got := getSecret()
		if got != secret {
			t.Errorf("getSecret() = %v, want %v", got, secret)
		}
	})

	t.Run("GetConfig error", func(t *testing.T) {
		secret = ""

		monkey.Patch(configs.GetConfig, func(key string) (string, error) {
			return "", errors.New("GetConfig error")
		})
		defer monkey.Unpatch(configs.GetConfig)

		got := getSecret()
		if got != "" {
			t.Errorf("getSecret() = %v, want %v", got, "")
		}
	})

	t.Run("get from config", func(t *testing.T) {
		secret = ""

		want := "123"
		monkey.Patch(configs.GetConfig, func(key string) (string, error) {
			return want, nil
		})
		defer monkey.Unpatch(configs.GetConfig)

		got := getSecret()
		if got != want {
			t.Errorf("getSecret() = %v, want %v", got, want)
		}
	})
}
