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

func Test_runFeedCardCmd(t *testing.T) {
	fakeExit := func(int) {
		log.Print("fake exit")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	const emptyMsg = "titles & picURLs & messageURLs can not be empty"
	const differentCountMsg = "titles & picURLs & messageURLs count must be equal"

	t.Run("titles is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		runFeedCardCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := emptyMsg

		if !strings.Contains(got, want) {
			t.Errorf("runFeedCardCmd() = %v, want %v", got, want)
		}
	})

	t.Run("titles & picURLs different count", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		feedCardVars.titles = []string{"1", "2"}
		feedCardVars.picURLs = []string{"1"}
		feedCardVars.messageURLs = []string{"1"}

		runFeedCardCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := differentCountMsg

		if !strings.Contains(got, want) {
			t.Errorf("runFeedCardCmd() = %v, want %v", got, want)
		}
	})

	t.Run("new client error", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		feedCardVars.titles = []string{"1"}
		feedCardVars.picURLs = []string{"1"}
		feedCardVars.messageURLs = []string{"1"}
		msg := "new client error"

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return nil, errors.New(msg)
		})
		defer monkey.Unpatch(newClient)

		runFeedCardCmd(&cobra.Command{}, []string{})

		got := buf.String()

		if !strings.Contains(got, msg) {
			t.Errorf("runFeedCardCmd() = %v, want %v", got, msg)
		}
	})

	t.Run("client send", func(t *testing.T) {
		feedCardVars.titles = []string{"1"}
		feedCardVars.picURLs = []string{"1"}
		feedCardVars.messageURLs = []string{"1"}
		client := &dingtalk.Client{}

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return client, nil
		})
		defer monkey.Unpatch(newClient)

		runFeedCardCmd(&cobra.Command{}, []string{})
	})
}
