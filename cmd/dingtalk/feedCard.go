package dingtalk

import (
	"github.com/CatchZeng/dingtalk"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var feedCardCmd = &cobra.Command{
	Use:   "feedCard",
	Short: "send feedCard message with Client robot",
	Long:  `send feedCard message with Client robot`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		if !CheckToken() {
			log.L(log.Red, "access_token can not be empty")
			return
		}

		if len(feedCardVars.titles) < 1 || len(feedCardVars.picURLs) < 1 || len(feedCardVars.messageURLs) < 1 {
			log.L(log.Red, "titles & picURLs & messageURLs can not be empty")
			return
		}

		if len(feedCardVars.titles) == len(feedCardVars.picURLs) && len(feedCardVars.picURLs) == len(feedCardVars.messageURLs) {
			client := dingtalk.NewClient(rootVars.accessToken, rootVars.secret)

			msg := dingtalk.NewFeedCardMessage()
			for i := 0; i < len(feedCardVars.titles); i++ {
				msg.AppendLink(feedCardVars.titles[i], feedCardVars.messageURLs[i], feedCardVars.picURLs[i])
			}
			if _, err := client.Send(msg); err != nil {
				log.L(log.Red, err.Error())
			}
		} else {
			log.L(log.Red, "titles & picURLs & messageURLs count must be equal")
			return
		}
	},
}

// FeedCardVars struct
type FeedCardVars struct {
	titles      []string
	picURLs     []string
	messageURLs []string
}

var feedCardVars FeedCardVars

func init() {
	rootCmd.AddCommand(feedCardCmd)

	feedCardCmd.Flags().StringSliceVarP(&feedCardVars.titles, "titles", "i", []string{}, "titles")
	feedCardCmd.Flags().StringSliceVarP(&feedCardVars.picURLs, "picURLs", "p", []string{}, "picURLs")
	feedCardCmd.Flags().StringSliceVarP(&feedCardVars.messageURLs, "messageURLs", "u", []string{}, "messageURLs")
}