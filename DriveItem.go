package go_microsoftgraph

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"time"
)

type DriveItem struct {
	DownloadUrl          string    `json:"@microsoft.graph.downloadUrl"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	ETag                 string    `json:"eTag"`
	Id                   string    `json:"id"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	Name                 string    `json:"name"`
	WebUrl               string    `json:"webUrl"`
	CTag                 string    `json:"cTag"`
	Size                 int       `json:"size"`
	CreatedBy            struct {
		User struct {
			Email       string `json:"email"`
			Id          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	LastModifiedBy struct {
		User struct {
			Email       string `json:"email"`
			Id          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"lastModifiedBy"`
	ParentReference struct {
		DriveType string `json:"driveType"`
		DriveId   string `json:"driveId"`
		Id        string `json:"id"`
		Name      string `json:"name"`
		Path      string `json:"path"`
		SiteId    string `json:"siteId"`
	} `json:"parentReference"`
	FileSystemInfo struct {
		CreatedDateTime      time.Time `json:"createdDateTime"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	} `json:"fileSystemInfo"`
	Folder *struct {
		ChildCount int `json:"childCount"`
	} `json:"folder"`
	Shared struct {
		Scope string `json:"scope"`
	} `json:"shared"`
}

type GetDriveItemConfig struct {
	DriveId     string
	DriveItemId string
}

func (service *Service) GetDriveItem(cfg *GetDriveItemConfig) (*DriveItem, *errortools.Error) {
	var driveItem DriveItem

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("drives/%s/items/%s", cfg.DriveId, cfg.DriveItemId)),
		ResponseModel: &driveItem,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &driveItem, nil
}

type ListDriveItemChildrenConfig struct {
	DriveId     string
	DriveItemId string
}

func (service *Service) ListDriveItemChildren(cfg *ListDriveItemChildrenConfig) (*[]DriveItem, *errortools.Error) {
	var response = struct {
		Value []DriveItem `json:"value"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("drives/%s/items/%s/children", cfg.DriveId, cfg.DriveItemId)),
		ResponseModel: &response,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.Value, nil
}
