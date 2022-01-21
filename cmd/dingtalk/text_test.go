package dingtalk

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"

	"bou.ke/monkey"
	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/spf13/cobra"
)

func Test_runTextCmd(t *testing.T) {
	fakeExit := func(int) {
		log.Print("fake exit")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	t.Run("content is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		runTextCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "content can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runTextCmd() = %v, want %v", got, want)
		}
	})

	t.Run("new client error", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		textVars.content = "123"
		msg := "new client error"

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return nil, errors.New(msg)
		})
		defer monkey.Unpatch(newClient)

		runTextCmd(&cobra.Command{}, []string{})

		got := buf.String()

		if !strings.Contains(got, msg) {
			t.Errorf("runTextCmd() = %v, want %v", got, msg)
		}
	})

	t.Run("client send", func(t *testing.T) {
		textVars.content = "123"
		client := &dingtalk.Client{}

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return client, nil
		})
		defer monkey.Unpatch(newClient)

		runTextCmd(&cobra.Command{}, []string{})
	})
}
