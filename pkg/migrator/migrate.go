package migrator

import (
	"github.com/Intware-Consulting/nexus-migrator/pkg/nexus"

	"github.com/rs/zerolog/log"
)

func Migrate(sourceRepository string, targetRepository string) error {
	log.Info().Str("command", "migrate").Str("source", sourceRepository).Str("target", targetRepository).Msgf("migrate repository [%s] -> [%s]", sourceRepository, targetRepository)

	nexusClient := nexus.CreateClient()
	assets, err := nexusClient.GetAssets(sourceRepository)
	if err != nil {
		log.Error().Err(err).Msgf("Error occurred when retrieving assets list from [%s] repository", sourceRepository)
		return err
	}

	errMigrate := nexusClient.MigrateAssets(assets, targetRepository)
	if errMigrate != nil {
		log.Error().Err(errMigrate).Msgf("Failed to migrate assets to [%s] repository", targetRepository)
		return errMigrate
	}
	return nil
}
