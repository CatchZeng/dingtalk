package cmd

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/CatchZeng/dingtalk/version"
	"github.com/spf13/cobra"
)

func Test_runVersionCmd(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	runVersionCmd(&cobra.Command{}, []string{})
	got := buf.String()

	want := version.GetVersion()

	if !strings.Contains(got, want) {
		t.Errorf("runVersionCmd() = %v, want %v", got, want)
	}
}
