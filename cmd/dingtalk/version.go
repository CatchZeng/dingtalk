package dingtalk

import (
	"log"

	v "github.com/CatchZeng/gutils/version"
	"github.com/spf13/cobra"
)

const (
	version   = "1.5.0"
	buildTime = "2022/04/20"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "dingtalk version",
	Long:  `dingtalk version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	v := v.Stringify(version, buildTime)
	log.Println(v)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
