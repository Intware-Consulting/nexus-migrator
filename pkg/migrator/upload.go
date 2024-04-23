package migrator

import (
	"fmt"

	"github.com/Intware-Consulting/nexus-migrator/pkg/nexus"
	"github.com/Intware-Consulting/nexus-migrator/pkg/utils"

	"github.com/rs/zerolog/log"
)

func Upload(sourceDirectory string, targetRepository string) error {
	log.Info().Str("command", "upload").Str("source", sourceDirectory).Str("target", targetRepository).Msgf("upload repository [%s] -> [%s]", sourceDirectory, targetRepository)

	list, err := utils.ListFilesInDirTree(sourceDirectory)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to list assets in directory [%s] repository", sourceDirectory)
		return err
	}

	if len(list) == 0 {
		return fmt.Errorf("no files in directory")
	}

	nexusClient := nexus.CreateClient()
	errUpload := nexusClient.UploadAssets(list, targetRepository)
	if errUpload != nil {
		log.Error().Err(errUpload).Msgf("Failed to upload assets to [%s] repository", targetRepository)
		return errUpload
	}
	return nil
}
