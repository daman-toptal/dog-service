package util

import (
	"dog-service/util/logging"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	breedApiUrlPattern = "https://dog.ceo/api/breed/%s/images/random/%d"
)

type breedApiMultipleResponse struct {
	Message []string `json:"message,omitempty"`
	Status  string   `json:"status,omitempty"`
}

func DownloadImage(imageURL string) (string, []byte, error) {

	// Build fileName from fullPath
	fileURL, err := url.Parse(imageURL)
	if err != nil {
		logging.Errorf("error parsing url - %v", err.Error())
		return "", nil, err
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Put content on byte slice
	resp, err := http.Get(imageURL)
	if err != nil {
		logging.Errorf("error getting image - %v", err.Error())
		return "", nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Errorf("error reading body - %v", err.Error())
	}

	return fileName, respBytes, err
}

func SaveImage(imageData []byte, fileName string) error {
	// Create blank file
	err := ioutil.WriteFile(fileName, imageData, 0644)
	return err
}

func GetImageURLs(breedName string, count int32) ([]string, error) {
	// Get Image URLs
	resp, err := http.Get(fmt.Sprintf(breedApiUrlPattern, breedName, count))
	if err != nil {
		logging.Errorf("error getting image urls - %v", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Errorf("error reading body - %v", err.Error())
		return nil, err
	}

	logging.Infof("raw response - %s", string(body))

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("breed not found")
	}

	apiResp := new(breedApiMultipleResponse)
	err = json.Unmarshal(body, apiResp)
	if err != nil {
		logging.Errorf("unmarshaling error - %v", err.Error())
		return nil, fmt.Errorf("fetch failed, received %d", resp.StatusCode)
	}

	if apiResp.Status != "success" || len(apiResp.Message) == 0 {
		logging.Errorf("invalid response - %v", string(body))
		return nil, fmt.Errorf("fetch failed, unexpected response")
	}

	return apiResp.Message, nil
}

func EnsureDir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
