package migrator

import (
	"github.com/Intware-Consulting/nexus-migrator/pkg/migrator"
	"github.com/Intware-Consulting/nexus-migrator/pkg/utils"
	"github.com/spf13/cobra"
)

// go run main.go download -s software -t /tmp
var downloadCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"m"},
	Short:   "Download repository.",
	Long: `Download remote repository to the local filesystem.

If files exists, the will be overwritten.`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceRepository, _ := cmd.Flags().GetString("source")
		targetDirectory, _ := cmd.Flags().GetString("target")

		cobra.CheckErr(utils.CheckIfDirExists(targetDirectory))

		err := migrator.Download(sourceRepository, targetDirectory)
		cobra.CheckErr(err)
	},
}

func init() {
	downloadCmd.PersistentFlags().StringP("source", "s", "", "Source repository name.")
	downloadCmd.MarkPersistentFlagRequired("source")
	downloadCmd.PersistentFlags().StringP("target", "t", "", "Target directory path.")
	downloadCmd.MarkPersistentFlagRequired("target")
}
