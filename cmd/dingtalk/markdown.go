package dingtalk

import (
	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "send markdown message with DingTalk robot",
	Long:  `send markdown message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runMarkdownCmd,
}

func runMarkdownCmd(_ *cobra.Command, args []string) {
	if len(markdownVars.title) < 1 {
		log.L(log.Red, "title can not be empty")
		return
	}

	if len(markdownVars.text) < 1 {
		log.L(log.Red, "text can not be empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.L(log.Red, err.Error())
		return
	}

	msg := dingtalk.NewMarkdownMessage().
		SetMarkdown(markdownVars.title, markdownVars.text).
		SetAt(atMobiles, isAtAll)
	req, _, err := client.Send(msg)
	if debug {
		log.L(log.Green, req)
	}
	if err != nil {
		log.L(log.Red, err.Error())
	}
}

// MarkdownVars struct
type MarkdownVars struct {
	title string
	text  string
}

var markdownVars MarkdownVars

func init() {
	rootCmd.AddCommand(markdownCmd)
	markdownCmd.Flags().StringVarP(&markdownVars.title, "title", "i", "", "title")
	markdownCmd.Flags().StringVarP(&markdownVars.text, "text", "e", "", "text")
}
