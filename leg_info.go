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

import (
	"time"
)

type LegInfo struct {

	// Action description of the call operation
	Action string `json:"action,omitempty"`

	// Call direction
	Direction string `json:"direction,omitempty"`

	// Call duration in seconds
	Duration int32 `json:"duration,omitempty"`

	// Information on extension
	Extension LegInfoExtensionInfo `json:"extension,omitempty"`

	// Leg type
	LegType string `json:"legType,omitempty"`

	// The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	StartTime time.Time `json:"startTime,omitempty"`

	// Call type
	Type_ string `json:"type,omitempty"`

	// Status description of the call operation
	Result string `json:"result,omitempty"`

	// Caller information
	From CallerInfo `json:"from,omitempty"`

	// Callee information
	To CallerInfo `json:"to,omitempty"`

	// Call transport
	Transport string `json:"transport,omitempty"`

	// Call recording data. Returned if the call is recorded
	Recording RecordingInfo `json:"recording,omitempty"`
}
