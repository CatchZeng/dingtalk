package dingtalk

import (
	"log"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/spf13/cobra"
)

var actionCardCmd = &cobra.Command{
	Use:   "actionCard",
	Short: "send actionCard message with DingTalk robot",
	Long:  `send actionCard message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runActionCardCmd,
}

func runActionCardCmd(_ *cobra.Command, args []string) {
	if len(actionCardVars.Title) < 1 {
		log.Fatal("title can not be empty")
		return
	}

	if len(actionCardVars.Text) < 1 {
		log.Fatal("text can not be empty")
		return
	}

	var isOverallJump = false
	if len(actionCardVars.SingleTitle) < 1 {
		if len(btnTitles) < 1 {
			log.Fatal("btns can not be empty when singleTitle is empty")
			return
		}
	} else {
		isOverallJump = true
		if len(actionCardVars.SingleURL) < 1 {
			log.Fatal("singleURL can not be empty")
			return
		}
	}

	client, err := newClient()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	msg := dingtalk.NewActionCardMessage()
	if isOverallJump {
		msg.SetOverallJump(
			actionCardVars.Title,
			actionCardVars.Text,
			actionCardVars.SingleTitle,
			actionCardVars.SingleURL,
			actionCardVars.BtnOrientation,
			actionCardVars.HideAvatar)
	} else {
		if len(btnTitles) != len(btnActionURLs) {
			log.Fatal("btnTitles & btnActionURLs count must be equal")
			return
		}

		for i := 0; i < len(btnTitles); i++ {
			actionCardVars.Btns = append(actionCardVars.Btns, dingtalk.Btn{
				Title:     btnTitles[i],
				ActionURL: btnActionURLs[i],
			})
		}
		msg.SetIndependentJump(
			actionCardVars.Title,
			actionCardVars.Text,
			actionCardVars.Btns,
			actionCardVars.BtnOrientation,
			actionCardVars.HideAvatar)
	}
	req, _, err := client.Send(msg)
	if debug {
		log.Print(req)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
}

var actionCardVars dingtalk.ActionCard
var btnTitles, btnActionURLs []string

func init() {
	rootCmd.AddCommand(actionCardCmd)
	actionCardCmd.Flags().StringVarP(&actionCardVars.Title, "title", "i", "", "title")
	actionCardCmd.Flags().StringVarP(&actionCardVars.Text, "text", "e", "", "text")
	actionCardCmd.Flags().StringVarP(&actionCardVars.SingleTitle, "singleTitle", "n", "", "singleTitle")
	actionCardCmd.Flags().StringVarP(&actionCardVars.SingleURL, "singleURL", "u", "", "singleURL")
	actionCardCmd.Flags().StringSliceVarP(&btnTitles, "btnTitles", "b", []string{}, "btnTitles")
	actionCardCmd.Flags().StringSliceVarP(&btnActionURLs, "btnActionURLs", "c", []string{}, "btnActionURLs")
	actionCardCmd.Flags().StringVarP(&actionCardVars.BtnOrientation, "btnOrientation", "o", "", "btnOrientation")
	actionCardCmd.Flags().StringVarP(&actionCardVars.HideAvatar, "hideAvatar", "d", "", "hideAvatar")
}
