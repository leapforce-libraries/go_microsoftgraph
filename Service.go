package go_microsoftgraph

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	go_token "github.com/leapforce-libraries/go_oauth2/token"
	"github.com/leapforce-libraries/go_oauth2/tokensource"
	"net/http"
	"time"
)

const (
	apiName            string = "MicrosoftGraph"
	apiUrl             string = "Me.go"
	authUrl            string = "https://login.microsoftonline.com/%s/oauth2/v2.0/authorize"
	tokenUrl           string = "https://login.microsoftonline.com/%s/oauth2/v2.0/token"
	tokenHttpMethod    string = http.MethodPost
	defaultRedirectUrl string = "http://localhost:8080/oauth/redirect"
	dateTimeLayout     string = "2006-01-02T15:04:05Z"
)

// Service stores GoogleService configuration
//
type Service struct {
	clientId      string
	clientSecret  string
	oAuth2Service *oauth2.Service
	errorResponse *ErrorResponse
}

type ServiceConfig struct {
	ClientId      string
	ClientSecret  string
	TenantId      string
	TokenSource   tokensource.TokenSource
	RedirectUrl   *string
	RefreshMargin *time.Duration
}

func NewService(cfg *ServiceConfig) (*Service, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if cfg.ClientId == "" {
		return nil, errortools.ErrorMessage("ClientId not provided")
	}

	if cfg.TenantId == "" {
		return nil, errortools.ErrorMessage("TenantId not provided")
	}

	redirectUrl := defaultRedirectUrl
	if cfg.RedirectUrl != nil {
		redirectUrl = *cfg.RedirectUrl
	}

	oauth2ServiceConfig := oauth2.ServiceConfig{
		ClientId:        cfg.ClientId,
		ClientSecret:    cfg.ClientSecret,
		RedirectUrl:     redirectUrl,
		AuthUrl:         fmt.Sprintf(authUrl, cfg.TenantId),
		TokenUrl:        fmt.Sprintf(tokenUrl, cfg.TenantId),
		RefreshMargin:   cfg.RefreshMargin,
		TokenHttpMethod: tokenHttpMethod,
		TokenSource:     cfg.TokenSource,
	}
	oauth2Service, e := oauth2.NewService(&oauth2ServiceConfig)
	if e != nil {
		return nil, e
	}

	return &Service{
		clientId:      cfg.ClientId,
		oAuth2Service: oauth2Service,
	}, nil
}

func (service *Service) HttpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	var request *http.Request
	var response *http.Response
	var e *errortools.Error

	// add error model
	service.errorResponse = &ErrorResponse{}
	requestConfig.ErrorModel = service.errorResponse

	request, response, e = service.oAuth2Service.HttpRequest(requestConfig)
	if e != nil {
		if service.errorResponse.Error.Message != "" {
			e.SetMessage(service.errorResponse.Error.Message)
		}
	}

	return request, response, e
}

func (service *Service) AuthorizeUrl(scope string) string {
	return service.oAuth2Service.AuthorizeUrl(scope, nil, nil, nil)
}

func (service *Service) ValidateToken() (*go_token.Token, *errortools.Error) {
	return service.oAuth2Service.ValidateToken()
}

func (service *Service) GetTokenFromCode(r *http.Request) *errortools.Error {
	return service.oAuth2Service.GetTokenFromCode(r, nil)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.clientId
}

func (service *Service) ApiCallCount() int64 {
	return service.oAuth2Service.ApiCallCount()
}

func (service *Service) ApiReset() {
	service.oAuth2Service.ApiReset()
}

func (service *Service) ErrorResponse() *ErrorResponse {
	return service.errorResponse
}
