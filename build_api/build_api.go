package buildapi

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type BuildAPI struct {
	BaseURL  string
	Username string
	APIKey   string
	AppID    string
}

type Package struct {
	Name         string  `json:"Name"`
	Status       string  `json:"Status"`
	Description  string  `json:"Description"`
	Creator      string  `json:"Creator"`
	CreationDate int     `json:"CreationDate"`
	Version      string  `json:"Version"`
	PackageID    string  `json:"PackageId"`
	Size         float64 `json:"Size"`
}

func (b BuildAPI) SetRequestHeaders(req http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Mendix-Username", b.Username)
	req.Header.Set("Mendix-ApiKey", b.APIKey)
}

func (b BuildAPI) RetrievePackages() []Package {
	client := http.Client{}
	req, _ := http.NewRequest("GET", b.BaseURL+"/apps/"+b.AppID+"/packages", nil)
	b.SetRequestHeaders(*req)
	response, err := client.Do(req)
	log.Info(response.Body)
	log.Info(err)
	jsonDataFromResp, _ := io.ReadAll(response.Body)

	var packages []Package
	if err := json.Unmarshal([]byte(jsonDataFromResp), &packages); err != nil {
		log.Error(err)
	}
	return packages
}

func (b BuildAPI) GetLatestPackage() Package {
	packages := b.RetrievePackages()
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].CreationDate > packages[j].CreationDate
	})

	return packages[0]
}

func (b BuildAPI) GetRevisionFromPackage(p Package) string {
	re := regexp.MustCompile(`.mda`)
	result := re.Split(p.Version, -1)
	revision := strings.SplitAfter(result[0], ".")
	return revision[len(revision)-1]
}
