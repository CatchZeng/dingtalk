package dingtalk

import (
	"github.com/CatchZeng/dingtalk"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "send text message with DingTalk robot",
	Long:  `send text message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		if len(textVars.content) < 1 {
			log.L(log.Red, "content can not be empty")
			return
		}

		client, err := newClient()
		if err != nil {
			log.L(log.Red, err.Error())
			return
		}

		msg := dingtalk.NewTextMessage().
			SetContent(textVars.content).
			SetAt(atMobiles, isAtAll)
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
