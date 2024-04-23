package nexus

import (
	"fmt"
	"path/filepath"

	"github.com/Intware-Consulting/nexus-migrator/pkg/utils"
	"github.com/spf13/viper"

	"github.com/rs/zerolog/log"
)

type Client struct {
	baseURL       string
	httpClient    *HTTPClient
	uriRestAPI    string
	uriRepository string
}

func CreateClient() *Client {
	baseURL := viper.GetString("nexus")
	username := viper.GetString("user")
	password := viper.GetString("pass")
	timeout := viper.GetDuration("timeout")

	httpClient := createHTTPClient(timeout, username, password)

	return &Client{
		baseURL:       baseURL,
		httpClient:    httpClient,
		uriRestAPI:    "/service/rest/v1",
		uriRepository: "/repository",
	}
}

func (n *Client) GetAssets(repository string) (*Assets, error) {
	url := fmt.Sprintf("%s%s/assets?repository=%s", n.baseURL, n.uriRestAPI, repository)
	totalAssets := &Assets{}

	finished := false
	for !finished {
		assets := &Assets{}
		err := n.httpClient.GetJSON(url, assets)
		if err != nil {
			return nil, err
		}

		log.Debug().Msgf("Found [%d] assets on url: [%s]", len(assets.Items), url)
		totalAssets.Items = append(totalAssets.Items, assets.Items...)
		if assets.ContinuationToken == nil {
			finished = true
		} else {
			url = fmt.Sprintf("%s%s/assets?repository=%s&continuationToken=%s", n.baseURL, n.uriRestAPI, repository, assets.ContinuationToken)
		}
	}
	log.Info().Msgf("Total assets found: %d", len(totalAssets.Items))

	return totalAssets, nil
}

func (n *Client) MigrateAssets(assets *Assets, targetRepository string) error {
	totalCount := len(assets.Items)
	for i, item := range assets.Items {
		targetURL := fmt.Sprintf("%s%s/%s/%s", n.baseURL, n.uriRepository, targetRepository, item.Path)
		log.Info().Msgf("[%d/%d] Migrating asset [%s] => [%s]", i+1, totalCount, item.DownloadURL, targetURL)

		err := n.httpClient.Copy(item.DownloadURL, targetURL)
		if err != nil {
			return err
		}
	}
	log.Info().Msg("Migration finished successfully")

	return nil
}

func (n *Client) DownloadAssets(assets *Assets, targetDirectoryPath string) error {
	totalCount := len(assets.Items)
	for i, item := range assets.Items {
		targetFilePath := fmt.Sprintf("%s/%s", targetDirectoryPath, item.Path)
		err := utils.EnsureDirectory(filepath.Dir(targetFilePath))
		if err != nil {
			return err
		}
		log.Info().Msgf("[%d/%d] Downloading asset [%s] => [%s]", i+1, totalCount, item.DownloadURL, targetFilePath)
		_, errDownload := n.httpClient.Download(item.DownloadURL, targetFilePath)
		if errDownload != nil {
			return errDownload
		}
		fmt.Println(item.Checksum.Md5)
		errChecksum := utils.VerifyChecksum(targetFilePath, item.Checksum.Md5, "md5")
		if errChecksum != nil {
			return errChecksum
		}
	}
	log.Info().Msg("Download finished successfully")

	return nil
}

func (n *Client) UploadAssets(assetsList map[string]string, targetRepository string) error {
	totalCount := len(assetsList)
	i := 1
	for filePath, repositoryPath := range assetsList {
		targetURL := fmt.Sprintf("%s%s/%s/%s", n.baseURL, n.uriRepository, targetRepository, repositoryPath)
		log.Info().Msgf("[%d/%d] Uploading asset [%s] => [%s]", i, totalCount, filePath, targetURL)

		err := n.httpClient.Upload(targetURL, filePath)
		if err != nil {
			return err
		}
		i++
	}
	log.Info().Msg("Upload finished successfully")
	return nil
}
