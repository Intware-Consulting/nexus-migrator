package migrator

import (
	"fmt"

	"github.com/Intware-Consulting/nexus-migrator/pkg/migrator"
	"github.com/Intware-Consulting/nexus-migrator/pkg/utils"
	"github.com/spf13/cobra"
)

// go run main.go upload -s /tmp -t software
var uploadCmd = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"m"},
	Short:   "Upload repositories",
	Long: `Upload files from local filesystem into remote repository.

If files exists, the will be overwritten.`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceDirectory, _ := cmd.Flags().GetString("source")
		targetRepository, _ := cmd.Flags().GetString("target")

		cobra.CheckErr(utils.CheckIfDirExists(sourceDirectory))

		fmt.Println(cmd.Parent().Flags().GetString("endpoint"))
		err := migrator.Upload(sourceDirectory, targetRepository)
		cobra.CheckErr(err)
	},
}

func init() {
	uploadCmd.PersistentFlags().StringP("source", "s", "", "Source directory path.")
	uploadCmd.MarkPersistentFlagRequired("source")
	uploadCmd.PersistentFlags().StringP("target", "t", "", "Target repository name.")
	uploadCmd.MarkPersistentFlagRequired("target")
}
