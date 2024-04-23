package migrator

import (
	"github.com/Intware-Consulting/nexus-migrator/pkg/migrator"
	"github.com/spf13/cobra"
)

// go run main.go migrate -s software -t test
var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"m"},
	Short:   "Migrate repositories.",
	Long: `Move files from one remote repository into another.
	
	If files exists, the will be overwritten.`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceRepository, _ := cmd.Flags().GetString("source")
		targetRepository, _ := cmd.Flags().GetString("target")
		err := migrator.Migrate(sourceRepository, targetRepository)
		cobra.CheckErr(err)
	},
}

func init() {
	migrateCmd.PersistentFlags().StringP("source", "s", "", "Source repository name.")
	migrateCmd.MarkPersistentFlagRequired("source")
	migrateCmd.PersistentFlags().StringP("target", "t", "", "Target repository name.")
	migrateCmd.MarkPersistentFlagRequired("target")
}
