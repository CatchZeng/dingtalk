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

func Test_runActionCardCmd(t *testing.T) {
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

		runActionCardCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "title can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runActionCardCmd() = %v, want %v", got, want)
		}
	})

	t.Run("text is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		actionCardVars.Title = "123"

		runActionCardCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "text can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runActionCardCmd() = %v, want %v", got, want)
		}
	})

	t.Run("singleTitle is empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		actionCardVars.Title = "123"
		actionCardVars.Text = "123"
		actionCardVars.SingleTitle = ""
		btnTitles = []string{}

		runActionCardCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "btns can not be empty when singleTitle is empty"

		if !strings.Contains(got, want) {
			t.Errorf("runActionCardCmd() = %v, want %v", got, want)
		}
	})

	t.Run("singleTitle is not empty", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		actionCardVars.Title = "123"
		actionCardVars.Text = "123"
		actionCardVars.SingleTitle = "123"
		actionCardVars.SingleURL = ""

		runActionCardCmd(&cobra.Command{}, []string{})
		got := buf.String()

		want := "singleURL can not be empty"

		if !strings.Contains(got, want) {
			t.Errorf("runActionCardCmd() = %v, want %v", got, want)
		}
	})

	t.Run("new client error", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		actionCardVars.Title = "123"
		actionCardVars.Text = "123"
		actionCardVars.SingleTitle = "123"
		actionCardVars.SingleURL = "123"
		msg := "new client error"

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return nil, errors.New(msg)
		})
		defer monkey.Unpatch(newClient)

		runActionCardCmd(&cobra.Command{}, []string{})

		got := buf.String()

		if !strings.Contains(got, msg) {
			t.Errorf("runActionCardCmd() = %v, want %v", got, msg)
		}
	})

	t.Run("btnTitles & btnActionURLs different count ", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		actionCardVars.Title = "123"
		actionCardVars.Text = "123"
		actionCardVars.SingleTitle = ""
		actionCardVars.SingleURL = ""
		btnTitles = []string{"1"}
		btnActionURLs = []string{"1", "2"}
		msg := "btnTitles & btnActionURLs count must be equal"

		client := &dingtalk.Client{}
		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return client, nil
		})
		defer monkey.Unpatch(newClient)

		runActionCardCmd(&cobra.Command{}, []string{})

		got := buf.String()

		if !strings.Contains(got, msg) {
			t.Errorf("runActionCardCmd() = %v, want %v", got, msg)
		}
	})

	t.Run("client send", func(t *testing.T) {
		actionCardVars.Title = "123"
		actionCardVars.Text = "123"
		actionCardVars.SingleTitle = "123"
		actionCardVars.SingleURL = "123"
		client := &dingtalk.Client{}

		monkey.Patch(newClient, func() (*dingtalk.Client, error) {
			return client, nil
		})
		defer monkey.Unpatch(newClient)

		runActionCardCmd(&cobra.Command{}, []string{})
	})
}
