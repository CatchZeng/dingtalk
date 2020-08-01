package dingtalk

import (
	"github.com/CatchZeng/dingtalk"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "send markdown message with Client robot",
	Long:  `send markdown message with Client robot`,
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

		client := dingtalk.NewClient(rootVars.accessToken, rootVars.secret)
		msg := dingtalk.NewMarkdownMessage().
			SetMarkdown(markdownVars.title, markdownVars.text).
			SetAt(rootVars.atMobiles, rootVars.isAtAll)
		if _, err := client.Send(msg); err != nil {
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
