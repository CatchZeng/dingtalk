package cmd

import (
	"github.com/CatchZeng/dingtalk/client"
	"github.com/CatchZeng/dingtalk/message"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "send text message with DingTalk robot",
	Long:  `send text message with DingTalk robot`,
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

		dingTalk := client.DingTalk{
			AccessToken: rootVars.accessToken,
			Secret:      rootVars.secret,
		}
		msg := message.NewTextMessage().
			SetContent(textVars.content).
			SetAt(rootVars.atMobiles, rootVars.isAtAll)
		if _, err := dingTalk.Send(msg); err != nil {
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
