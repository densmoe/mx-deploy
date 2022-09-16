package deployapi

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type DeployAPI struct {
	BaseURL  string
	Username string
	APIKey   string
}

type App struct {
	Name      string `json:"Name"`
	URL       string `json:"Url"`
	ProjectId string `json:"ProjectId"`
	AppId     string `json:"AppId"`
}

func (d DeployAPI) RetrieveApps() []App {
	client := http.Client{}
	req, _ := http.NewRequest("GET", d.BaseURL+"/apps", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Mendix-Username", d.Username)
	req.Header.Set("Mendix-ApiKey", d.APIKey)
	response, _ := client.Do(req)
	jsonDataFromResp, _ := io.ReadAll(response.Body)

	var apps []App
	if err := json.Unmarshal([]byte(jsonDataFromResp), &apps); err != nil {
		log.Error(err)
	}
	return apps
}

func (d DeployAPI) RetrieveApp(appId string) App {
	client := http.Client{}
	req, _ := http.NewRequest("GET", d.BaseURL+"/apps/"+appId, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Mendix-Username", d.Username)
	req.Header.Set("Mendix-ApiKey", d.APIKey)
	response, _ := client.Do(req)
	jsonDataFromResp, _ := io.ReadAll(response.Body)

	var app App
	if err := json.Unmarshal([]byte(jsonDataFromResp), &app); err != nil {
		log.Error(err)
	}
	return app
}
