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

type InlineResponseDefault4 struct {

	// List of extensions belonging to a given department
	Records []DepartmentResponseExtensionInfo `json:"records,omitempty"`

	// Information on navigation
	Navigation NavigationInfo `json:"navigation,omitempty"`

	// Information on paging
	Paging PagingInfo `json:"paging,omitempty"`
}
