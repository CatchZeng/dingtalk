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

func Test_runLinkCmd(t *testing.T) {
	fakeExit := func(int) {
		log.Print("fake exit")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	t.Run("title is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		runLinkCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "title can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runLinkCmd() = %v, want %v", got, want)
		}
	})

	t.Run("text is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		linkVars.title = "123"

		runLinkCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "text can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runLinkCmd() = %v, want %v", got, want)
		}
	})

	t.Run("messageURL is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		linkVars.title = "123"
		linkVars.text = "123"

		runLinkCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "messageURL can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runLinkCmd() = %v, want %v", got, want)
		}
	})

	t.Run("new client error", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		linkVars.title = "123"
		linkVars.text = "123"
		linkVars.messageURL = "123"
		msg := "new client error"

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return nil, errors.New(msg)
		})
		defer monkey.Unpatch(newClient)

		runLinkCmd(&cobra.Command{}, []string{})

		got := buf.String()

		if !strings.Contains(got, msg) {
			t.Errorf("runLinkCmd() = %v, want %v", got, msg)
		}
	})

	t.Run("client send", func(t *testing.T) {
		linkVars.title = "123"
		linkVars.text = "123"
		linkVars.messageURL = "123"
		client := &dingtalk.Client{}

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return client, nil
		})
		defer monkey.Unpatch(newClient)

		runLinkCmd(&cobra.Command{}, []string{})
	})
}
