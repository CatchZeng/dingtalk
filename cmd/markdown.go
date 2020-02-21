package cmd

import (
	"github.com/CatchZeng/dingtalk/client"
	"github.com/CatchZeng/dingtalk/message"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "send markdown message with DingTalk robot",
	Long:  `send markdown message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		if !CheckToken() {
			log.L(log.Red, "access_token can not be empty")
			return
		}

		if len(markdownVars.title) < 1 {
			log.L(log.Red, "title can not be empty")
			return
		}

		if len(markdownVars.text) < 1 {
			log.L(log.Red, "text can not be empty")
			return
		}

		dingTalk := client.DingTalk{
			AccessToken: rootVars.accessToken,
			Secret:      rootVars.secret,
		}
		msg := message.NewMarkdownMessage().
			SetMarkdown(markdownVars.title, markdownVars.text).
			SetAt(rootVars.atMobiles, rootVars.isAtAll)
		if _, err := dingTalk.Send(msg); err != nil {
			log.L(log.Red, err.Error())
		}
	},
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
