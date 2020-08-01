package dingtalk

import (
	"github.com/CatchZeng/dingtalk/internal/version"
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
	v := version.GetVersion()
	log.Println(v)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
