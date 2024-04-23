package migrator

import (
	"github.com/Intware-Consulting/nexus-migrator/pkg/nexus"

	"github.com/rs/zerolog/log"
)

func Download(sourceRepository string, targetDirectory string) error {
	log.Info().Str("command", "download").Str("source", sourceRepository).Str("target", targetDirectory).Msgf("download repository [%s] -> [%s]", sourceRepository, targetDirectory)

	nexusClient := nexus.CreateClient()
	assets, err := nexusClient.GetAssets(sourceRepository)
	if err != nil {
		log.Error().Err(err).Msgf("Error occurred when retrieving assets list from [%s] repository", sourceRepository)
		return err
	}

	errDownload := nexusClient.DownloadAssets(assets, targetDirectory)
	if errDownload != nil {
		log.Error().Err(errDownload).Msgf("Failed to download assets to [%s] repository", targetDirectory)
		return errDownload
	}
	return nil
}
