package deployapiv4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type DeployAPIv4 struct {
	PAT string
}

type Member struct {
	UserID string `json:"userId"`
}

// PermissionSet represents the structure of the permission set
type PermissionSet struct {
	Member              Member `json:"member"`
	CanDeployApp        string `json:"canDeployApp"`
	CanManageBackups    string `json:"canManageBackups"`
	CanViewAlerts       string `json:"canViewAlerts"`
	CanAccessAPI        string `json:"canAccessAPI"`
	CanViewLogs         string `json:"canViewLogs"`
	CanManagePrivileges string `json:"canManagePrivileges"`
}

// ApiResponse represents the structure of the JSON response
type ApiResponse struct {
	TechnicalContact string          `json:"technicalContact"`
	Permissions      []PermissionSet `json:"permissions"`
	Pagination       struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Size   int `json:"size"`
	} `json:"pagination"`
}

// PatchPermissionsBody represents the structure of the PATCH request body
type PatchPermissionsBody struct {
	Permissions []PermissionSet `json:"permissions"`
}

// PatchTechnicalContactBody represents the structure of the PATCH request body {
type PatchTechnicalContactBody struct {
	TechnicalContact struct {
		UserID string `json:"userId"`
	} `json:"technicalContact"`
}

// AppResponse represents the structure of the JSON response
type GetAppsResponse struct {
	Apps       []App `json:"apps"`
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Size   int `json:"size"`
	} `json:"pagination"`
}

// App represents the structure of an individual app in the response
type App struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	LicenseType      string           `json:"licenseType"`
	Subdomain        string           `json:"subdomain"`
	TechnicalContact TechnicalContact `json:"technicalContact"`
	Region           string           `json:"region"`
}

// Package represents the package information in the response
type Package struct {
	ID             string    `json:"id"`
	AppID          string    `json:"appId"`
	ModelVersion   string    `json:"modelVersion"`
	RuntimeVersion string    `json:"runtimeVersion"`
	CreatedOn      time.Time `json:"createdOn"`
	ExpiresOn      time.Time `json:"expiresOn"`
	Description    string    `json:"description"`
	FileName       string    `json:"fileName"`
	FileSize       int       `json:"fileSize"`
}

// Environment represents an environment in the response
type Environment struct {
	ID           string  `json:"id"`
	AppID        string  `json:"appId"`
	Name         string  `json:"name"`
	State        string  `json:"state"`
	IsProduction bool    `json:"isProduction"`
	URL          string  `json:"url"`
	DBVersion    string  `json:"dbVersion"`
	PlanName     string  `json:"planName"`
	Package      Package `json:"package"`
}

// Response represents the overall response from the endpoint
type GetEnvironmentsResponse struct {
	Environments []Environment `json:"environments"`
	Pagination   struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Size   int `json:"size"`
	} `json:"pagination"`
}

// TechnicalContact represents the structure of the technical contact information
type TechnicalContact struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

const BaseURL = "https://cloud.home.mendix.com/api/v4"

func (d DeployAPIv4) SetRequestHeaders(req http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("MxToken %s", d.PAT))
}

func (d DeployAPIv4) GetLicensedApps() []App {
	client := http.Client{}

	apps := []App{}
	hasMorePages := true
	offset := 0
	for hasMorePages {
		url := fmt.Sprintf("%s/apps?licenseType=licensed&offset=%d", BaseURL, offset)
		req, _ := http.NewRequest("GET", url, nil)
		d.SetRequestHeaders(*req)
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return []App{}
		}
		defer response.Body.Close()
		fmt.Println(response.Status)

		// Read the response body
		body, err := io.ReadAll(response.Body)
		fmt.Println(string(body))
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return []App{}
		}

		// Parse the JSON response into the ApiResponse struct
		var apiResponse GetAppsResponse
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return []App{}
		}
		apps = append(apps, apiResponse.Apps...)
		hasMorePages = apiResponse.Pagination.Size == apiResponse.Pagination.Limit
		offset = offset + apiResponse.Pagination.Size
	}

	return apps
}

func (d DeployAPIv4) GetEnvironments(appId string) []Environment {
	client := http.Client{}
	url := fmt.Sprintf("%s/apps/%s/environments?limit=100", BaseURL, appId)
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	d.SetRequestHeaders(*req)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return []Environment{}
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []Environment{}
	}

	// Parse the JSON response into the ApiResponse struct
	var apiResponse GetEnvironmentsResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return []Environment{}
	}
	return apiResponse.Environments
}

func (d DeployAPIv4) GetUserPermissionsForEnvironment(appId string, environmentId string, userId string) ApiResponse {
	client := http.Client{}
	url := fmt.Sprintf("%s/apps/%s/environments/%s/permissions?userId=%s", BaseURL, appId, environmentId, userId)
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	d.SetRequestHeaders(*req)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ApiResponse{}
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ApiResponse{}
	}

	// Parse the JSON response into the ApiResponse struct
	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return ApiResponse{}
	}
	return apiResponse
}

func (d DeployAPIv4) SetUserPermissionsForEnvironment(appId string, environmentId string, userId string, canDeployApp bool, canManageBackups bool, canViewAlerts bool, canAccessAPI bool, canViewLogs bool, canManagePrivileges bool) PatchPermissionsBody {
	client := http.Client{}
	url := fmt.Sprintf("%s/apps/%s/environments/%s/permissions", BaseURL, appId, environmentId)
	fmt.Println(url)
	// Create the PATCH request body
	patchRequestBody := PatchPermissionsBody{
		Permissions: []PermissionSet{
			{
				Member:              Member{UserID: userId},
				CanDeployApp:        strconv.FormatBool(canDeployApp),
				CanManageBackups:    strconv.FormatBool(canManageBackups),
				CanViewAlerts:       strconv.FormatBool(canViewAlerts),
				CanAccessAPI:        strconv.FormatBool(canAccessAPI),
				CanViewLogs:         strconv.FormatBool(canViewLogs),
				CanManagePrivileges: strconv.FormatBool(canManagePrivileges),
			},
		},
	}

	// Convert the PATCH request body to JSON
	patchRequestBodyJSON, err := json.Marshal(patchRequestBody)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return PatchPermissionsBody{}
	}

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(patchRequestBodyJSON))
	d.SetRequestHeaders(*req)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return PatchPermissionsBody{}
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return PatchPermissionsBody{}
	}

	// Parse the JSON response into the ApiResponse struct
	var apiResponse PatchPermissionsBody
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return PatchPermissionsBody{}
	}
	return apiResponse
}

func (d DeployAPIv4) PatchTechnicalContact(appId string, userId string) {
	client := http.Client{}
	url := fmt.Sprintf("%s/apps/%s", BaseURL, appId)
	fmt.Println(url)
	// Create the PATCH request body
	patchRequestBody := PatchTechnicalContactBody{
		TechnicalContact: struct {
			UserID string `json:"userId"`
		}{
			UserID: userId,
		},
	}

	// Convert the PATCH request body to JSON
	patchRequestBodyJSON, err := json.Marshal(patchRequestBody)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(patchRequestBodyJSON))
	d.SetRequestHeaders(*req)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()
	fmt.Println(response.Status)
}
