package dingtalk

import (
	"log"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
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
		log.Fatal("title can not be empty")
	}

	if len(markdownVars.text) < 1 {
		log.Fatal("text can not be empty")
	}

	client, err := newClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	msg := dingtalk.NewMarkdownMessage().
		SetMarkdown(markdownVars.title, markdownVars.text).
		SetAt(atMobiles, isAtAll)
	req, _, err := client.Send(msg)
	if debug {
		log.Print(req)
	}
	if err != nil {
		log.Fatal(err.Error())
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
