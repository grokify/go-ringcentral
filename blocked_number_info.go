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

type BlockedNumberInfo struct {

	// Standard resource properties ID and canonical URI, see the section called “Resource Identification Properties”
	Id string `json:"id,omitempty"`

	// Canonical URI of a blocked number resource
	Uri string `json:"uri,omitempty"`

	// Name assigned by a user to a blocked phone number
	Name string `json:"name,omitempty"`

	// Phone number to be blocked
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
