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

type Body3 struct {

	// Company business name
	Company string `json:"company,omitempty"`

	// Company business email address
	Email string `json:"email,omitempty"`

	// Company business address
	BusinessAddress BusinessAddressInfo `json:"businessAddress,omitempty"`
}
