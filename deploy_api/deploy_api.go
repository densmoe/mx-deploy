package deployapi

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type DeployAPI struct {
	BaseURL   string
	Username  string
	APIKey    string
	AppID     string
	ProjectID string
}

type App struct {
	Name      string `json:"Name"`
	URL       string `json:"Url"`
	ProjectId string `json:"ProjectId"`
	AppId     string `json:"AppId"`
}

type Environment struct {
	Status        string `json:"Status"`
	EnvironmentId string `json:"EnvironmentId"`
	Mode          string `json:"Mode"`
	Url           string `json:"Url"`
	ModelVersion  string `json:"ModelVersion"`
	MendixVersion string `json:"MendixVersion"`
	IsProduction  bool   `json:"Production"`
}

func (d DeployAPI) SetRequestHeaders(req http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Mendix-Username", d.Username)
	req.Header.Set("Mendix-ApiKey", d.APIKey)
}

func (d DeployAPI) RetrieveApps() []App {
	client := http.Client{}
	req, _ := http.NewRequest("GET", d.BaseURL+"/apps", nil)
	d.SetRequestHeaders(*req)
	response, _ := client.Do(req)
	jsonDataFromResp, _ := io.ReadAll(response.Body)

	var apps []App
	if err := json.Unmarshal([]byte(jsonDataFromResp), &apps); err != nil {
		log.Error(err)
	}
	return apps
}

func (d DeployAPI) GetAppIdForProjectId(projectId string) string {
	var appId string
	apps := d.RetrieveApps()

	for _, app := range apps {

		if app.ProjectId == d.ProjectID {
			appId = app.AppId
			break
		}
	}

	return appId
}

func (d *DeployAPI) SetAppIdForProjectId(projectId string) {
	// This sets the DeployAPI.AppID for a given ProjectID
	appId := d.GetAppIdForProjectId(d.ProjectID)
	d.AppID = appId
}

func (d DeployAPI) RetrieveApp() App {
	client := http.Client{}
	req, _ := http.NewRequest("GET", d.BaseURL+"/apps/"+d.AppID, nil)
	d.SetRequestHeaders(*req)
	response, err := client.Do(req)
	log.Info(response.Body)
	log.Info(err)
	jsonDataFromResp, _ := io.ReadAll(response.Body)

	var app App
	if err := json.Unmarshal([]byte(jsonDataFromResp), &app); err != nil {
		log.Error(err)
	}
	return app
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

func (d DeployAPI) RetrieveEnvironments(appId string) []Environment {
	client := http.Client{}
	req, _ := http.NewRequest("GET", d.BaseURL+"/apps/"+appId+"/environments", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Mendix-Username", d.Username)
	req.Header.Set("Mendix-ApiKey", d.APIKey)
	response, _ := client.Do(req)
	jsonDataFromResp, _ := io.ReadAll(response.Body)

	var environments []Environment
	if err := json.Unmarshal([]byte(jsonDataFromResp), &environments); err != nil {
		log.Error(err)
	}
	return environments
}
