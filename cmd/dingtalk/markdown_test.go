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

func Test_runMarkdownCmd(t *testing.T) {
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

		runMarkdownCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "title can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runMarkdownCmd() = %v, want %v", got, want)
		}
	})

	t.Run("text is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		markdownVars.title = "123"

		runMarkdownCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "text can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runMarkdownCmd() = %v, want %v", got, want)
		}
	})

	t.Run("new client error", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		markdownVars.title = "123"
		markdownVars.text = "123"
		msg := "new client error"

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return nil, errors.New(msg)
		})
		defer monkey.Unpatch(newClient)

		runMarkdownCmd(&cobra.Command{}, []string{})

		got := buf.String()

		if !strings.Contains(got, msg) {
			t.Errorf("runMarkdownCmd() = %v, want %v", got, msg)
		}
	})

	t.Run("client send", func(t *testing.T) {
		markdownVars.title = "123"
		markdownVars.text = "123"
		client := &dingtalk.Client{}

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return client, nil
		})
		defer monkey.Unpatch(newClient)

		runMarkdownCmd(&cobra.Command{}, []string{})
	})
}
