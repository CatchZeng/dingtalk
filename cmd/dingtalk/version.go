package dingtalk

import (
	"github.com/CatchZeng/gutils/version"
	"log"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "dingtalk version",
	Long:  `dingtalk version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	v := version.Stringify("2.1.0", "2020/08/19")
	log.Println(v)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
