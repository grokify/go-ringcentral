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

type GroupInfo struct {

	// Internal identifier of a group
	Id string `json:"id,omitempty"`

	// Canonical URI of a group
	Uri string `json:"uri,omitempty"`

	// Amount of contacts in a group
	ContactsCount int32 `json:"contactsCount,omitempty"`

	// Name of a group
	GroupName string `json:"groupName,omitempty"`

	// Notes for a group
	Notes string `json:"notes,omitempty"`
}
