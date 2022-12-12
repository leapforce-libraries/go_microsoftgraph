package go_microsoftgraph

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type BookingStaffMember struct {
	Id                                       string        `json:"id"`
	DisplayName                              string        `json:"displayName"`
	EmailAddress                             string        `json:"emailAddress"`
	AvailabilityIsAffectedByPersonalCalendar bool          `json:"availabilityIsAffectedByPersonalCalendar"`
	Role                                     string        `json:"role"`
	UseBusinessHours                         bool          `json:"useBusinessHours"`
	SsEmailNotificationEnabled               bool          `json:"isEmailNotificationEnabled"`
	TimeZone                                 string        `json:"timeZone"`
	WorkingHours                             []WorkingHour `json:"workingHours"`
}

type WorkingHour struct {
	Day       string                `json:"day"`
	TimeSlots []WorkingHourTimeSlot `json:"timeSlots"`
}

type WorkingHourTimeSlot struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type GetBookingStaffMemberConfig struct {
	BookingBusinessId    string
	BookingStaffMemberId string
}

func (service *Service) GetBookingStaffMember(cfg *GetBookingStaffMemberConfig) (*BookingStaffMember, *errortools.Error) {
	var bookingStaffMember BookingStaffMember

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("solutions/bookingBusinesses/%s/staffMembers/%s", cfg.BookingBusinessId, cfg.BookingStaffMemberId)),
		ResponseModel: &bookingStaffMember,
	}
	_, _, e := service.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &bookingStaffMember, nil
}
