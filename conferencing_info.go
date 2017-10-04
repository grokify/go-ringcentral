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

type ConferencingInfo struct {

	// Canonical URI of a conferencing
	Uri string `json:"uri,omitempty"`

	// Determines if host user allows conference participants to join before the host
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost,omitempty"`

	// Access code for a host user
	HostCode string `json:"hostCode,omitempty"`

	// Internal parameter specifying conferencing engine
	Mode string `json:"mode,omitempty"`

	// Access code for any participant
	ParticipantCode string `json:"participantCode,omitempty"`

	// Primary conference phone number for user's home country returned in E.164 (11-digits) format
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Short URL leading to the service web page Tap to Join for audio conference bridge
	TapToJoinUri string `json:"tapToJoinUri,omitempty"`

	// List of multiple dial-in phone numbers to connect to audio conference service, relevant for user's brand. Each number is given with the country and location information, in order to let the user choose the less expensive way to connect to a conference. The first number in the list is the primary conference number, that is default and domestic
	PhoneNumbers []ConferencingInfoPhoneNumberInfo `json:"phoneNumbers,omitempty"`
}
