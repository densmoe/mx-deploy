package deployapiv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

type DeployAPIv2 struct {
	Username  string
	APIKey    string
	AppID     string
	ProjectID string
}

type UploadResponse struct {
	PackageId string
	JobId     string
}

type StatusResponse struct {
	Status string
}

const BaseURL = "https://deploy.mendix.com/api/v2"

func (d DeployAPIv2) UploadPackage(appId string, filePath string) UploadResponse {

	// Extract the file name from the filePath
	fileName := filepath.Base(filePath)

	var requestBody bytes.Buffer

	// Create a new multipart writer
	multipartWriter := multipart.NewWriter(&requestBody)

	// Add the file to the request
	fileWriter, err := multipartWriter.CreateFormFile("file", filePath)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		panic(err)
	}

	// Add headers
	multipartWriter.WriteField("Mendix-Username", d.Username)
	multipartWriter.WriteField("Mendix-ApiKey", d.APIKey)

	// Close the multipart writer
	multipartWriter.Close()

	url := fmt.Sprintf(`%s/apps/%s/packages/upload?name=%s`, BaseURL, appId, fileName)
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		panic(err)
	}

	// Set the Content-Type header to match the multipart form data
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	request.Header.Set("Mendix-Username", d.Username)
	request.Header.Set("Mendix-ApiKey", d.APIKey)

	// Perform the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Print the response
	jsonDataFromResp, _ := io.ReadAll(response.Body)
	var uploadResponse UploadResponse
	if err := json.Unmarshal([]byte(jsonDataFromResp), &uploadResponse); err != nil {
		log.Error(err)
	}
	return uploadResponse
}

func (d DeployAPIv2) GetUploadStatus(appId string, jobId string) StatusResponse {
	client := http.Client{}
	url := fmt.Sprintf(`%s/apps/%s/jobs/%s`, BaseURL, appId, jobId)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Mendix-Username", d.Username)
	request.Header.Set("Mendix-ApiKey", d.APIKey)
	response, err := client.Do(request)
	if err != nil {
		log.Error(err)
		panic("error")
	}
	jsonDataFromResp, _ := io.ReadAll(response.Body)

	var statusResponse StatusResponse
	if err := json.Unmarshal([]byte(jsonDataFromResp), &statusResponse); err != nil {
		log.Error(err)
	}
	return statusResponse
}

func (d DeployAPIv2) PollForUploadStatus(appId string, jobId string, timeout time.Duration) StatusResponse {

	latestStatusResponse := StatusResponse{}
	timeoutChan := time.After(timeout)

	for {
		select {
		case <-timeoutChan:
			log.Error("Upload timed out")
			return latestStatusResponse
		default:
			latestStatusResponse = d.GetUploadStatus(appId, jobId)

			switch latestStatusResponse.Status {
			case "Running":
			case "Queued":
			case "Completed":
				return latestStatusResponse
			case "Failed":
				return latestStatusResponse
			default:
				panic("panic!")
			}

			// Sleep for a while before making the next request
			time.Sleep(1 * time.Second)
		}
	}
}
