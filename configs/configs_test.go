package configs

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	t.Run("homedir.Dir() return error", func(t *testing.T) {
		monkey.Patch(homedir.Dir, func() (string, error) {
			return "", errors.New("homedir error")
		})
		shouldPanic(t, InitConfig)
	})

	t.Run("viper.ReadInConfig() return error", func(t *testing.T) {
		monkey.Patch(homedir.Dir, func() (string, error) {
			return "/catchzeng", nil
		})

		monkey.Patch(viper.ReadInConfig, func() error {
			return errors.New("ReadInConfig error")
		})

		InitConfig()
	})

	t.Run("viper.ReadInConfig() return nil", func(t *testing.T) {
		monkey.Patch(homedir.Dir, func() (string, error) {
			return "/catchzeng", nil
		})

		monkey.Patch(viper.ReadInConfig, func() error {
			return nil
		})

		InitConfig()
	})
}

func shouldPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should have panicked")
		}
	}()
	f()
}

func TestGetConfig(t *testing.T) {
	t.Run("viper.ReadInConfig() return error", func(t *testing.T) {
		monkey.Patch(viper.ReadInConfig, func() error {
			return errors.New("ReadInConfig error")
		})

		key := "123"
		if _, err := GetConfig(key); err == nil {
			t.Error("GetConfig error")
		}
	})

	t.Run("viper.ReadInConfig() return nil", func(t *testing.T) {
		monkey.Patch(viper.ReadInConfig, func() error {
			return nil
		})

		monkey.Patch(viper.GetString, func(key string) string {
			return key
		})

		key := "456"
		if value, err := GetConfig(key); err != nil || value != key {
			t.Error("GetConfig error")
		}
	})
}
