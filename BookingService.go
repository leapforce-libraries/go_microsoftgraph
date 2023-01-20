package go_microsoftgraph

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
	"time"
)

type BookingService struct {
	AdditionalInformation   *string         `json:"additionalInformation"`
	CustomQuestions         json.RawMessage `json:"customQuestions"`
	DefaultDuration         *string         `json:"defaultDuration"`
	DefaultLocation         *Location       `json:"defaultLocation"`
	DefaultPrice            *float64        `json:"defaultPrice"`
	DefaultPriceType        *string         `json:"defaultPriceType"`
	DefaultReminders        json.RawMessage `json:"defaultReminders"`
	Description             *string         `json:"description	"`
	DisplayName             string          `json:"displayName"`
	Id                      string          `json:"id"`
	IsAnonymousJoinEnabled  bool            `json:"isAnonymousJoinEnabled"`
	IsHiddenFromCustomers   bool            `json:"isHiddenFromCustomers"`
	IsLocationOnline        bool            `json:"isLocationOnline"`
	MaximumAttendeesCount   *int32          `json:"maximumAttendeesCount"`
	Notes                   *string         `json:"notes"`
	PostBuffer              *string         `json:"postBuffer"`
	PreBuffer               *string         `json:"preBuffer"`
	SchedulingPolicy        json.RawMessage `json:"schedulingPolicy"`
	SmsNotificationsEnabled bool            `json:"smsNotificationsEnabled"`
	StaffMemberIds          *[]string       `json:"staffMemberIds"`
	WebUrl                  *string         `json:"webUrl"`
}

type GetBookingServiceConfig struct {
	BookingBusinessId string
	BookingServiceId  string
}

func (service *Service) GetBookingService(cfg *GetBookingServiceConfig) (*BookingService, *errortools.Error) {
	var bookingService BookingService

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("solutions/bookingBusinesses/%s/services/%s", cfg.BookingBusinessId, cfg.BookingServiceId)),
		ResponseModel: &bookingService,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &bookingService, nil
}

type ListBookingServicesConfig struct {
	BookingBusinessId string
	Filter            *string
	Select            *string
}

func (service *Service) ListBookingServices(cfg *ListBookingServicesConfig) (*[]BookingService, *errortools.Error) {
	var values = url.Values{}
	if cfg.Filter != nil {
		values.Set("$filter", *cfg.Filter)
	}
	if cfg.Select != nil {
		values.Set("$select", *cfg.Select)
	}

	var response = struct {
		Value []BookingService `json:"value"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("solutions/bookingBusinesses/%s/services?%s", cfg.BookingBusinessId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.Value, nil
}

type ListCalendarViewBookingServicesConfig struct {
	BookingBusinessId string
	Start             time.Time
	End               time.Time
	Filter            *string
	Select            *string
}

func (service *Service) ListCalendarViewBookingServices(cfg *ListCalendarViewBookingServicesConfig) (*[]BookingService, *errortools.Error) {
	var values = url.Values{}
	values.Set("start", cfg.Start.Format(dateTimeLayout))
	values.Set("end", cfg.End.Format(dateTimeLayout))
	if cfg.Filter != nil {
		values.Set("$filter", *cfg.Filter)
	}
	if cfg.Select != nil {
		values.Set("$select", *cfg.Select)
	}

	var response = struct {
		Value []BookingService `json:"value"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("solutions/bookingBusinesses/%s/calendarView?%s", cfg.BookingBusinessId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.Value, nil
}
