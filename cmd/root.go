package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dingtalk",
	Short: "dingtalk is a command line tool for DingTalk",
	Long:  "dingtalk is a command line tool for DingTalk",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// CheckToken check token
func CheckToken() bool {
	return len(rootVars.accessToken) > 0
}

// RootVars struct
type RootVars struct {
	accessToken string
	secret      string
	isAtAll     bool
	atMobiles   []string
}

var rootVars RootVars

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootVars.accessToken, "token", "t", "", "access_token")
	rootCmd.PersistentFlags().StringVarP(&rootVars.secret, "secret", "s", "", "secret")
	rootCmd.PersistentFlags().BoolVarP(&rootVars.isAtAll, "isAtAll", "a", false, "isAtAll")
	rootCmd.PersistentFlags().StringArrayVarP(&rootVars.atMobiles, "atMobiles", "m", []string{}, "atMobiles")
}
