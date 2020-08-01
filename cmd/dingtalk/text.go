package dingtalk

import (
	"github.com/CatchZeng/dingtalk"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "send text message with Client robot",
	Long:  `send text message with Client robot`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		if !CheckToken() {
			log.L(log.Red, "access_token can not be empty")
			return
		}

		if len(textVars.content) < 1 {
			log.L(log.Red, "content can not be empty")
			return
		}

		client := dingtalk.NewClient(rootVars.accessToken, rootVars.secret)
		msg := dingtalk.NewTextMessage().
			SetContent(textVars.content).
			SetAt(rootVars.atMobiles, rootVars.isAtAll)
		if _, err := client.Send(msg); err != nil {
			log.L(log.Red, err.Error())
		}
	},
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
