package dingtalk

import (
	"bytes"
	"github.com/CatchZeng/gutils/version"
	"log"
	"os"
	"strings"
	"testing"

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

	want := version.Stringify("2.1.0", "2020/08/19")

	if !strings.Contains(got, want) {
		t.Errorf("runVersionCmd() = %v, want %v", got, want)
	}
}
