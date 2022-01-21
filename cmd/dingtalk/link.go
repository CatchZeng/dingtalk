package dingtalk

import (
	"log"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "send link message with DingTalk robot",
	Long:  `send link message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runLinkCmd,
}

func runLinkCmd(_ *cobra.Command, args []string) {
	if len(linkVars.title) < 1 {
		log.Fatal("title can not be empty")
		return
	}

	if len(linkVars.text) < 1 {
		log.Fatal("text can not be empty")
		return
	}

	if len(linkVars.messageURL) < 1 {
		log.Fatal("messageURL can not be empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	msg := dingtalk.NewLinkMessage().
		SetLink(linkVars.title, linkVars.text, linkVars.picURL, linkVars.messageURL)
	req, _, err := client.Send(msg)
	if debug {
		log.Print(req)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
}

// LinkVars struct
type LinkVars struct {
	title      string
	text       string
	picURL     string
	messageURL string
}

var linkVars LinkVars

func init() {
	rootCmd.AddCommand(linkCmd)
	linkCmd.Flags().StringVarP(&linkVars.title, "title", "i", "", "title")
	linkCmd.Flags().StringVarP(&linkVars.text, "text", "e", "", "text")
	linkCmd.Flags().StringVarP(&linkVars.picURL, "picURL", "p", "", "picURL")
	linkCmd.Flags().StringVarP(&linkVars.messageURL, "messageURL", "u", "", "messageURL")
}
