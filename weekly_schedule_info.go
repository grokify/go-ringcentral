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

type WeeklyScheduleInfo struct {

	// Time intervals for a particular day
	Monday []TimeInterval `json:"monday,omitempty"`

	// Time intervals for a particular day
	Tuesday []TimeInterval `json:"tuesday,omitempty"`

	// Time intervals for a particular day
	Wednesday []TimeInterval `json:"wednesday,omitempty"`

	// Time intervals for a particular day
	Thursday []TimeInterval `json:"thursday,omitempty"`

	// Time intervals for a particular day
	Friday []TimeInterval `json:"friday,omitempty"`

	// Time intervals for a particular day
	Saturday []TimeInterval `json:"saturday,omitempty"`

	// Time intervals for a particular day
	Sunday []TimeInterval `json:"sunday,omitempty"`
}
