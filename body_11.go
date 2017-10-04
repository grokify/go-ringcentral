/*
 * Ringcentral API
 *
 * RingCentral Connect Platform API
 *
 * OpenAPI spec version: v1.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ringcentral

type Body11 struct {

	// Topic of a meeting
	Topic string `json:"topic,omitempty"`

	// Type of a meeting. 'Instant' - meeting that is instantly started as soon as the host creates it; 'Scheduled' - common scheduled meeting; 'Recurring' - a recurring meeting. If the specified meeting type is 'Scheduled' then schedule property is mandatory for request
	MeetingType string `json:"meetingType,omitempty"`

	// Password required to join a meeting. Max number of characters is 10
	Password string `json:"password,omitempty"`

	// Schedule of a meeting
	Schedule MeetingScheduleInfo `json:"schedule,omitempty"`

	// If 'True' then the meeting can be joined and started by any participant (not host only). Supported for the meetings of 'Scheduled' and 'Recurring' type.
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost,omitempty"`

	// Enables starting video when host joins the meeting
	StartHostVideo bool `json:"startHostVideo,omitempty"`

	// Enables starting video when participants join the meeting
	StartParticipantsVideo bool `json:"startParticipantsVideo,omitempty"`

	// Meeting audio options. Possible values are 'Phone', 'ComputerAudio'
	AudioOptions []string `json:"audioOptions,omitempty"`
}
