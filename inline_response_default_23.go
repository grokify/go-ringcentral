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

type InlineResponseDefault23 struct {

	// List of phone numbers
	Records []PhoneNumberInfo `json:"records,omitempty"`

	// Information on navigation
	Navigation NavigationInfo `json:"navigation,omitempty"`

	// Information on paging
	Paging PagingInfo `json:"paging,omitempty"`
}
