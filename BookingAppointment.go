package go_microsoftgraph

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
	"time"
)

type BookingAppointment struct {
	Id                       string     `json:"id"`
	SelfServiceAppointmentId string     `json:"selfServiceAppointmentId"`
	IsLocationOnline         bool       `json:"isLocationOnline"`
	JoinWebUrl               *string    `json:"joinWebUrl"`
	Customers                []Customer `json:"customers"`
	CustomerTimeZone         *string    `json:"customerTimeZone"`
	SmsNotificationsEnabled  bool       `json:"smsNotificationsEnabled"`
	ServiceId                *string    `json:"serviceId"`
	ServiceName              *string    `json:"serviceName"`
	Duration                 *string    `json:"duration"`
	PreBuffer                *string    `json:"preBuffer"`
	PostBuffer               *string    `json:"postBuffer"`
	PriceType                *string    `json:"priceType"`
	Price                    *float64   `json:"price"`
	ServiceNotes             *string    `json:"serviceNotes"`
	OptOutOfCustomerEmail    bool       `json:"optOutOfCustomerEmail"`
	StaffMemberIds           []string   `json:"staffMemberIds"`
	StartDateTime            *DateTime  `json:"startDateTime"`
	EndDateTime              *DateTime  `json:"endDateTime"`
	ServiceLocation          Location   `json:"serviceLocation"`
	Reminders                []string   `json:"reminders"`
}

type Customer struct {
	CustomerId   string   `json:"customerId"`
	Name         string   `json:"name"`
	EmailAddress *string  `json:"emailAddress"`
	Phone        *string  `json:"phone"`
	Notes        *string  `json:"notes"`
	Location     Location `json:"location"`
	TimeZone     *string  `json:"timeZone"`
}

type Location struct {
	DisplayName          string  `json:"displayName"`
	LocationEmailAddress *string `json:"locationEmailAddress"`
	LocationUri          *string `json:"locationUri"`
	LocationType         *string `json:"locationType"`
	UniqueId             *string `json:"uniqueId"`
	UniqueIdType         *string `json:"uniqueIdType"`
	Address              struct {
		Street          string `json:"street"`
		City            string `json:"city"`
		State           string `json:"state"`
		CountryOrRegion string `json:"countryOrRegion"`
		PostalCode      string `json:"postalCode"`
	} `json:"address"`
	Coordinates struct {
		Altitude         float64 `json:"altitude"`
		Latitude         float64 `json:"latitude"`
		Longitude        float64 `json:"longitude"`
		Accuracy         float64 `json:"accuracy"`
		AltitudeAccuracy float64 `json:"altitudeAccuracy"`
	} `json:"coordinates"`
}

type DateTime struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type ListBookingAppointmentsConfig struct {
	BookingBusinessId string
}

func (service *Service) ListBookingAppointments(cfg *ListBookingAppointmentsConfig) (*[]BookingAppointment, *errortools.Error) {
	var response = struct {
		Value []BookingAppointment `json:"value"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("solutions/bookingBusinesses/%s/appointments", cfg.BookingBusinessId)),
		ResponseModel: &response,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.Value, nil
}

type ListCalendarViewBookingAppointmentsConfig struct {
	BookingBusinessId string
	Start             time.Time
	End               time.Time
}

func (service *Service) ListCalendarViewBookingAppointments(cfg *ListCalendarViewBookingAppointmentsConfig) (*[]BookingAppointment, *errortools.Error) {
	var values = url.Values{}
	values.Set("start", cfg.Start.Format(dateTimeLayout))
	values.Set("end", cfg.End.Format(dateTimeLayout))

	var response = struct {
		Value []BookingAppointment `json:"value"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("solutions/bookingBusinesses/%s/calendarView?%s", cfg.BookingBusinessId, values.Encode())),
		ResponseModel: &response,
	}
	fmt.Println(requestConfig.Url)
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.Value, nil
}
