package go_microsoftgraph

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type Me struct {
	Id                string   `json:"id"`
	BusinessPhones    []string `json:"businessPhones"`
	DisplayName       *string  `json:"displayName"`
	GivenName         *string  `json:"givenName"`
	JobTitle          *string  `json:"jobTitle"`
	Mail              *string  `json:"mail"`
	MobilePhone       *string  `json:"mobilePhone"`
	OfficeLocation    *string  `json:"officeLocation"`
	PreferredLanguage *string  `json:"preferredLanguage"`
	Surname           *string  `json:"surname"`
	UserPrincipalName *string  `json:"userPrincipalName"`
}

func (service *Service) Me() (*Me, *errortools.Error) {
	var me Me

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("me"),
		ResponseModel: &me,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &me, nil
}
