package nexus

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type HTTPClient struct {
	http        *http.Client
	credentials Credentials
}

type Credentials struct {
	username string
	password string
}

func (c *Credentials) addBasicAuth(req *http.Request) {
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.username+":"+c.password)))
}

func createHTTPClient(timeout time.Duration, username string, password string) *HTTPClient {
	client := &http.Client{
		Timeout: timeout,
	}

	creds := Credentials{
		username: username,
		password: password,
	}

	return &HTTPClient{
		http:        client,
		credentials: creds,
	}
}

func (c *HTTPClient) GetJSON(url string, response interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	c.credentials.addBasicAuth(req)
	req.Header.Add("accept", "application/json")
	res, err := c.http.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("[%d] %s", res.StatusCode, res.Status)
	}

	return json.NewDecoder(res.Body).Decode(response)
}

func (c *HTTPClient) Copy(sourceURL string, targetURL string) error {
	reqDownload, _ := http.NewRequest(http.MethodGet, sourceURL, nil)
	c.credentials.addBasicAuth(reqDownload)

	resDownload, err := c.http.Do(reqDownload)
	if err != nil {
		return err
	}
	defer resDownload.Body.Close()

	if resDownload.StatusCode != http.StatusOK {
		return fmt.Errorf("[%d] %s", resDownload.StatusCode, resDownload.Status)
	}

	reqUpload, _ := http.NewRequest(http.MethodPut, targetURL, resDownload.Body)
	c.credentials.addBasicAuth(reqUpload)

	resUpload, err := c.http.Do(reqUpload)
	if err != nil {
		return err
	}

	if resUpload.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("[%d] %s", resUpload.StatusCode, resUpload.Status)
	}
	defer resUpload.Body.Close()

	return nil
}

func (c *HTTPClient) Download(sourceURL string, targetFilePath string) (int64, error) {
	reqDownload, _ := http.NewRequest(http.MethodGet, sourceURL, nil)
	c.credentials.addBasicAuth(reqDownload)

	resDownload, err := c.http.Do(reqDownload)
	if err != nil {
		return 0, err
	}
	defer resDownload.Body.Close()

	if resDownload.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("[%d] %s", resDownload.StatusCode, resDownload.Status)
	}

	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		return 0, err
	}
	defer targetFile.Close()

	return io.Copy(targetFile, resDownload.Body)
}

func (c *HTTPClient) Upload(targetURL string, sourceFilePath string) error {
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}

	reqUpload, _ := http.NewRequest(http.MethodPut, targetURL, sourceFile)
	c.credentials.addBasicAuth(reqUpload)

	resUpload, err := c.http.Do(reqUpload)
	if err != nil {
		return err
	}

	if resUpload.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("[%d] %s", resUpload.StatusCode, resUpload.Status)
	}
	defer resUpload.Body.Close()

	return nil
}
