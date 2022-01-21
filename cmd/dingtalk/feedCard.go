package dingtalk

import (
	"log"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/spf13/cobra"
)

var feedCardCmd = &cobra.Command{
	Use:   "feedCard",
	Short: "send feedCard message with DingTalk robot",
	Long:  `send feedCard message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runFeedCardCmd,
}

func runFeedCardCmd(_ *cobra.Command, args []string) {
	if len(feedCardVars.titles) < 1 || len(feedCardVars.picURLs) < 1 || len(feedCardVars.messageURLs) < 1 {
		log.Fatal("titles & picURLs & messageURLs can not be empty")
		return
	}

	if len(feedCardVars.titles) == len(feedCardVars.picURLs) && len(feedCardVars.picURLs) == len(feedCardVars.messageURLs) {
		client, err := newClient()
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		msg := dingtalk.NewFeedCardMessage()
		for i := 0; i < len(feedCardVars.titles); i++ {
			msg.AppendLink(feedCardVars.titles[i], feedCardVars.messageURLs[i], feedCardVars.picURLs[i])
		}
		req, _, err := client.Send(msg)
		if debug {
			log.Print(req)
		}
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	} else {
		log.Fatal("titles & picURLs & messageURLs count must be equal")
	}
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
