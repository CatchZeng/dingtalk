package dingtalk

import (
	"log"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "send text message with DingTalk robot",
	Long:  `send text message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runTextCmd,
}

func runTextCmd(_ *cobra.Command, _ []string) {
	if len(textVars.content) < 1 {
		log.Fatal("content can not be empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	msg := dingtalk.NewTextMessage().
		SetContent(textVars.content).
		SetAt(atMobiles, isAtAll)
	req, _, err := client.Send(msg)
	if debug {
		log.Print(req)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
}

// TextVars struct
type TextVars struct {
	content string
}

var textVars TextVars

func init() {
	rootCmd.AddCommand(textCmd)
	textCmd.Flags().StringVarP(&textVars.content, "content", "c", "", "content")
}
