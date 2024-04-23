package nexus

import "time"

type Assets struct {
	Items []struct {
		DownloadURL string `json:"downloadUrl"`
		Path        string `json:"path"`
		ID          string `json:"id"`
		Repository  string `json:"repository"`
		Format      string `json:"format"`
		Checksum    struct {
			Sha1   string `json:"sha1"`
			Sha512 string `json:"sha512"`
			Sha256 string `json:"sha256"`
			Md5    string `json:"md5"`
		} `json:"checksum"`
		ContentType    string    `json:"contentType"`
		LastModified   time.Time `json:"lastModified"`
		LastDownloaded time.Time `json:"lastDownloaded"`
		Uploader       string    `json:"uploader"`
		UploaderIP     string    `json:"uploaderIp"`
		FileSize       int       `json:"fileSize"`
		BlobCreated    time.Time `json:"blobCreated"`
	} `json:"items"`
	ContinuationToken interface{} `json:"continuationToken"`
}
