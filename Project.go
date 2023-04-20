package googleresourcemanager

import (
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type ProjectsResponse struct {
	Projects      []Project `json:"projects"`
	NextPageToken string    `json:"nextPageToken"`
}

type Project struct {
	ProjectNumber  string    `json:"projectNumber"`
	ProjectId      string    `json:"projectId"`
	LifecycleState string    `json:"lifecycleState"`
	Name           string    `json:"name"`
	CreateTime     time.Time `json:"createTime"`
	Parent         struct {
		Type string `json:"type"`
		Id   string `json:"id"`
	} `json:"parent"`
}

func (service *Service) Projects() (*[]Project, *errortools.Error) {
	var projects []Project

	var pageToken = ""

	for {
		var path = "projects"
		if pageToken != "" {
			path = fmt.Sprintf("%s?pageToken=%s", path, pageToken)
		}

		var projectsResponse ProjectsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(path),
			ResponseModel: &projectsResponse,
		}

		_, _, e := service.googleService.HttpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		projects = append(projects, projectsResponse.Projects...)

		if projectsResponse.NextPageToken == "" {
			break
		}

		pageToken = projectsResponse.NextPageToken
	}

	return &projects, nil
}
